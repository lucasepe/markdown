// Package url provides functions for parsing, decoding and encoding URLs.
package url

// A URL represents a parsed URL.
type URL struct {
	Scheme      string
	RawScheme   string
	Auth        string
	Host        string
	Port        string
	Path        string
	RawQuery    string
	Fragment    string
	HasQuery    bool
	HasFragment bool
	IPv6        bool
	Slashes     bool
}

// String reassembles the URL into a URL string.
func (u *URL) String() string {
	size := len(u.Path)
	if u.Scheme != "" {
		size += len(u.Scheme) + 1
	}
	if u.Slashes {
		size += 2
	}
	if u.Auth != "" {
		size += len(u.Auth) + 1
	}
	if u.Host != "" {
		size += len(u.Host)
		if u.IPv6 {
			size += 2
		}
	}
	if u.Port != "" {
		size += len(u.Port) + 1
	}
	if u.HasQuery {
		size += len(u.RawQuery) + 1
	}
	if u.HasFragment {
		size += len(u.Fragment) + 1
	}
	if size == 0 {
		return ""
	}

	buf := make([]byte, size)
	i := 0
	if u.Scheme != "" {
		i += copy(buf, u.Scheme)
		buf[i] = ':'
		i++
	}
	if u.Slashes {
		buf[i] = '/'
		i++
		buf[i] = '/'
		i++
	}
	if u.Auth != "" {
		i += copy(buf[i:], u.Auth)
		buf[i] = '@'
		i++
	}
	if u.Host != "" {
		if u.IPv6 {
			buf[i] = '['
			i++
			i += copy(buf[i:], u.Host)
			buf[i] = ']'
			i++
		} else {
			i += copy(buf[i:], u.Host)
		}
	}
	if u.Port != "" {
		buf[i] = ':'
		i++
		i += copy(buf[i:], u.Port)
	}
	i += copy(buf[i:], u.Path)
	if u.HasQuery {
		buf[i] = '?'
		i++
		i += copy(buf[i:], u.RawQuery)
	}
	if u.HasFragment {
		buf[i] = '#'
		i++
		copy(buf[i:], u.Fragment)
	}
	return string(buf)
}
