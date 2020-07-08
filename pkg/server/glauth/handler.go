package glauth

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/glauth/glauth/pkg/config"
	"github.com/glauth/glauth/pkg/handler"
	"github.com/glauth/glauth/pkg/stats"
	ber "github.com/nmcclain/asn1-ber"
	"github.com/nmcclain/ldap"
	accounts "github.com/owncloud/ocis-accounts/pkg/proto/v0"
	"github.com/owncloud/ocis-pkg/v2/log"
)

type queryType string

const (
	usersQuery  queryType = "users"
	groupsQuery queryType = "groups"
)

type ocisHandler struct {
	as  accounts.AccountsService
	log log.Logger
	cfg *config.Config
}

func (h ocisHandler) Bind(bindDN, bindSimplePw string, conn net.Conn) (ldap.LDAPResultCode, error) {
	bindDN = strings.ToLower(bindDN)
	baseDN := strings.ToLower("," + h.cfg.Backend.BaseDN)

	h.log.Debug().
		Str("binddn", bindDN).
		Str("basedn", h.cfg.Backend.BaseDN).
		Interface("src", conn.RemoteAddr()).
		Msg("Bind request")

	stats.Frontend.Add("bind_reqs", 1)

	// parse the bindDN - ensure that the bindDN ends with the BaseDN
	if !strings.HasSuffix(bindDN, baseDN) {
		h.log.Error().
			Str("binddn", bindDN).
			Str("basedn", h.cfg.Backend.BaseDN).
			Interface("src", conn.RemoteAddr()).
			Msg("BindDN not part of our BaseDN")
		return ldap.LDAPResultInvalidCredentials, nil
	}
	parts := strings.Split(strings.TrimSuffix(bindDN, baseDN), ",")
	if len(parts) > 2 {
		h.log.Error().
			Str("binddn", bindDN).
			Int("numparts", len(parts)).
			Interface("src", conn.RemoteAddr()).
			Msg("BindDN should have only one or two parts")
		return ldap.LDAPResultInvalidCredentials, nil
	}
	userName := strings.TrimPrefix(parts[0], "cn=")

	// check password
	_, err := h.as.ListAccounts(context.TODO(), &accounts.ListAccountsRequest{
		//Query: fmt.Sprintf("username eq '%s'", username),
		// TODO this allows lookung up users when you know the username using basic auth
		// adding the password to the query is an option but sending the sover the wira a la scim seems ugly
		// but to set passwords our accounts need it anyway
		Query: fmt.Sprintf("login eq '%s' and password eq '%s'", userName, bindSimplePw),
	})
	if err != nil {
		h.log.Error().
			Str("username", userName).
			Str("binddn", bindDN).
			Interface("src", conn.RemoteAddr()).
			Msg("Login failed")
		return ldap.LDAPResultInvalidCredentials, nil
	}

	stats.Frontend.Add("bind_successes", 1)
	h.log.Debug().
		Str("binddn", bindDN).
		Interface("src", conn.RemoteAddr()).
		Msg("Bind success")
	return ldap.LDAPResultSuccess, nil
}

func (h ocisHandler) Search(bindDN string, searchReq ldap.SearchRequest, conn net.Conn) (ldap.ServerSearchResult, error) {
	bindDN = strings.ToLower(bindDN)
	baseDN := strings.ToLower("," + h.cfg.Backend.BaseDN)
	searchBaseDN := strings.ToLower(searchReq.BaseDN)
	h.log.Debug().
		Str("binddn", bindDN).
		Str("basedn", h.cfg.Backend.BaseDN).
		Str("filter", searchReq.Filter).
		Interface("src", conn.RemoteAddr()).
		Msg("Search request")
	stats.Frontend.Add("search_reqs", 1)

	// validate the user is authenticated and has appropriate access
	if len(bindDN) < 1 {
		return ldap.ServerSearchResult{
			ResultCode: ldap.LDAPResultInsufficientAccessRights,
		}, fmt.Errorf("search error: Anonymous BindDN not allowed %s", bindDN)
	}
	if !strings.HasSuffix(bindDN, baseDN) {
		return ldap.ServerSearchResult{
			ResultCode: ldap.LDAPResultInsufficientAccessRights,
		}, fmt.Errorf("search error: BindDN %s not in our BaseDN %s", bindDN, h.cfg.Backend.BaseDN)
	}
	if !strings.HasSuffix(searchBaseDN, h.cfg.Backend.BaseDN) {
		return ldap.ServerSearchResult{
			ResultCode: ldap.LDAPResultInsufficientAccessRights,
		}, fmt.Errorf("search error: search BaseDN %s is not in our BaseDN %s", searchBaseDN, h.cfg.Backend.BaseDN)
	}

	var qtype queryType = ""
	query := ""
	var err error
	if searchReq.Filter == "(&)" { // see Absolute True and False Filters in https://tools.ietf.org/html/rfc4526#section-2
		query = ""
	} else {
		var cf *ber.Packet
		cf, err = ldap.CompileFilter(searchReq.Filter)
		if err != nil {
			return ldap.ServerSearchResult{
				ResultCode: ldap.LDAPResultOperationsError,
			}, fmt.Errorf("Search Error: error parsing filter: %s", searchReq.Filter)
		}
		qtype, query, err = parseFilter(cf)
		if err != nil {
			return ldap.ServerSearchResult{
				ResultCode: ldap.LDAPResultOperationsError,
			}, fmt.Errorf("Search Error: error parsing filter: %s", searchReq.Filter)
		}
	}

	entries := []*ldap.Entry{}
	h.log.Debug().
		Str("binddn", bindDN).
		Str("basedn", h.cfg.Backend.BaseDN).
		Str("filter", searchReq.Filter).
		Str("qtype", string(qtype)).
		Str("query", query).
		Msg("parsed query")
	switch qtype {
	case usersQuery:
		accounts, err := h.as.ListAccounts(context.TODO(), &accounts.ListAccountsRequest{
			Query: query,
		})
		if err != nil {
			h.log.Error().
				Err(err).
				Str("binddn", bindDN).
				Str("basedn", h.cfg.Backend.BaseDN).
				Str("filter", searchReq.Filter).
				Str("query", query).
				Interface("src", conn.RemoteAddr()).
				Msg("Could not list accounts")

			return ldap.ServerSearchResult{
				ResultCode: ldap.LDAPResultOperationsError,
			}, errors.New("search error: error getting users")
		}
		entries = append(entries, h.mapAccounts(accounts.Accounts)...)
	}

	stats.Frontend.Add("search_successes", 1)
	h.log.Debug().
		Str("binddn", bindDN).
		Str("basedn", h.cfg.Backend.BaseDN).
		Str("filter", searchReq.Filter).
		Interface("src", conn.RemoteAddr()).
		Msg("AP: Search OK")

	return ldap.ServerSearchResult{
		Entries:    entries,
		Referrals:  []string{},
		Controls:   []ldap.Control{},
		ResultCode: ldap.LDAPResultSuccess,
	}, nil
}

func attribute(name string, values ...string) *ldap.EntryAttribute {
	return &ldap.EntryAttribute{
		Name:   name,
		Values: values,
	}
}

func (h ocisHandler) mapAccounts(accounts []*accounts.Account) []*ldap.Entry {
	var entries []*ldap.Entry
	for _, acc := range accounts {
		attrs := []*ldap.EntryAttribute{
			attribute("objectClass", "posixAccount", "inetOrgPerson", "organizationalPerson", "Person", "top"),
			attribute("cn", acc.PreferredName),
			attribute("uid", acc.PreferredName),
			attribute("sn", acc.PreferredName),
		}
		if acc.DisplayName != "" {
			attrs = append(attrs, attribute("displayName", acc.DisplayName))
		}
		if acc.Mail != "" {
			attrs = append(attrs, attribute("mail", acc.Mail))
		}
		if acc.UidNumber != 0 { // TODO no root?
			attrs = append(attrs, attribute("uidnumber", strconv.FormatInt(acc.UidNumber, 10)))
		}
		if acc.GidNumber != 0 {
			attrs = append(attrs, attribute("gidnumber", strconv.FormatInt(acc.GidNumber, 10)))
		}
		if acc.Description != "" {
			attrs = append(attrs, attribute("description", acc.Description))
		}

		dn := fmt.Sprintf("%s=%s,%s=%s,%s",
			h.cfg.Backend.NameFormat,
			acc.PreferredName,
			h.cfg.Backend.GroupFormat,
			"users",
			h.cfg.Backend.BaseDN,
		)
		entries = append(entries, &ldap.Entry{DN: dn, Attributes: attrs})
	}
	return entries
}

// LDAP filters might ask for grouips and users at the same time, eg.
// (|
//   (&(objectClass=posixaccount)(cn=einstein))
//   (&(objectClass=posixgroup)(cn=users))
// )

// (&(objectClass=posixaccount)(objectClass=posixgroup))
// qtype is one of
// "" not determined
// "users"
// "groups"
func parseFilter(f *ber.Packet) (qtype queryType, q string, err error) {
	switch ldap.FilterMap[f.Tag] {
	case "Equality Match":
		if len(f.Children) != 2 {
			return "", "", errors.New("equality match must have only two children")
		}
		attribute := strings.ToLower(f.Children[0].Value.(string))
		value := f.Children[1].Value.(string)

		// replace attributes
		switch attribute {
		case "objectclass":
			switch value {
			case "posixaccount", "shadowaccount", "users", "person", "inetorgperson", "organizationalperson":
				qtype = usersQuery
			case "posixgroup", "groups":
				qtype = groupsQuery
			default:
				qtype = ""
			}
		case "cn", "uid":
			q = fmt.Sprintf("preferred_name eq '%s'", escapeValue(value))
		case "mail":
			q = fmt.Sprintf("mail eq '%s'", escapeValue(value))
		case "displayname":
			q = fmt.Sprintf("display_name eq '%s'", escapeValue(value))
		case "uidnumber":
			q = fmt.Sprintf("uid_number eq '%s'", escapeValue(value))
		case "gidnumber":
			q = fmt.Sprintf("gid_number eq '%s'", escapeValue(value))
		case "description":
			q = fmt.Sprintf("description eq '%s'", escapeValue(value))
		default:
			err = fmt.Errorf("filter by %s not implmented", attribute)
		}

		return
	case "And":
		subQueries := []string{}
		for _, child := range f.Children {
			var subQuery string
			var qt queryType
			qt, subQuery, err = parseFilter(child)
			if err != nil {
				return "", "", err
			}
			if qtype == "" {
				qtype = qt
			} else if qt != "" && qt != qtype {
				return "", "", fmt.Errorf("mixing user and group filters not supported")
			}
			if subQuery != "" {
				subQueries = append(subQueries, subQuery)
			}
		}
		return qtype, strings.Join(subQueries, " and "), nil
	case "Or":
		subQueries := []string{}
		for _, child := range f.Children {
			var subQuery string
			var qt queryType
			qt, subQuery, err = parseFilter(child)
			if err != nil {
				return "", "", err
			}
			if qtype == "" {
				qtype = qt
			} else if qt != "" && qt != qtype {
				return "", "", fmt.Errorf("mixing user and group filters not supported")
			}
			if subQuery != "" {
				subQueries = append(subQueries, subQuery)
			}
		}
		return qtype, strings.Join(subQueries, " or "), nil
	case "Not":
		if len(f.Children) != 1 {
			return "", "", errors.New("not filter must have only one child")
		}
		qtype, subQuery, err := parseFilter(f.Children[0])
		if err != nil {
			return "", "", err
		}
		if subQuery != "" {
			q = fmt.Sprintf("not %s", subQuery)
		}
		return qtype, q, nil
	}
	return
}

// escapeValue escapes all special characters in the value
func escapeValue(value string) string {
	return strings.ReplaceAll(value, "'", "''")
}

func (h ocisHandler) Close(boundDN string, conn net.Conn) error {
	stats.Frontend.Add("closes", 1)
	return nil
}

// NewOCISHandler implements a glauth backend with ocis-accounts as tdhe datasource
func NewOCISHandler(opts ...Option) handler.Handler {
	options := newOptions(opts...)

	handler := ocisHandler{
		log: options.Logger,
		cfg: options.Config,
		as:  options.AccountsService,
	}
	return handler
}
