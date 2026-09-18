package handlers

import (
	"log"
	"net/http"

	"github.com/coder/websocket"

	"github.com/gen2brain/cam2ip/image"
)

// Socket handler.
type Socket struct {
	stream *Stream
}

// NewSocket returns new socket handler.
func NewSocket(stream *Stream) *Socket {
	return &Socket{stream}
}

// ServeHTTP handles requests on incoming connections.
func (s *Socket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("socket: accept: %v", err)

		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "")

	// Binary clients receive raw JPEG frames; text clients receive base64 for backward compatibility.
	binary := r.URL.Query().Get("binary") == "1"

	// CloseRead handles control frames and cancels ctx when the client goes away.
	ctx := conn.CloseRead(r.Context())

	ch := s.stream.subscribe()
	defer s.stream.unsubscribe(ch)

	for {
		select {
		case <-ctx.Done():
			return

		case frame := <-ch:
			var err error
			if binary {
				err = conn.Write(ctx, websocket.MessageBinary, frame)
			} else {
				err = conn.Write(ctx, websocket.MessageText, []byte(image.EncodeToString(frame)))
			}

			if err != nil {
				return
			}
		}
	}
}
