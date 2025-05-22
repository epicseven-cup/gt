package commit

type Header struct {
	headers []string
	head    int
	tail    int
}

func NewHeader(headers []string) *Header {
	return &Header{
		headers: headers,
		head:    0,
		tail:    len(headers),
	}
}

// parseCache reads the cache files for recently used headers and stores them into the Header struct
func parseCache() (string, error) {
	return "", nil
}

// AppendHeader adds new header into the head of the array and shift
func (h *Header) AppendHeader(v string) string {
	h.headers[h.head] = v
	h.head = (h.head + 1) % len(h.headers)
	h.tail = (h.tail + 1) % len(h.headers)
	return v
}

func (h *Header) Head() string {
	return h.headers[h.head]
}
func (h *Header) Tail() string {
	return h.headers[h.tail]
}

func (h *Header) Len() int {
	return len(h.headers)
}

func (h *Header) List() []string {
	return h.headers
}

func (h *Header) Clear() {
	h.headers = []string{}
}

func (h *Header) ReadCache(configPath string) error {
	if configPath == "" {
		return nil
	}

}
