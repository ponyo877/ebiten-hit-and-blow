# Ebiten Hit & Blow

Welcome to **Ebiten Hit & Blow**, a peer-to-peer implementation of the classic Hit & Blow game using WebRTC and Go WebAssembly using Ebitengine. This project allows you to match and play with random strangers in real-time.

## Demo

Check out the live demo: [Hit & Blow](https://hit-and-blow.pages.dev/go/)

<img width="371" alt="image" src="https://github.com/user-attachments/assets/4be793e1-2b1f-4daf-bfe4-c5012c725bba">

## Features

- **Real-time Multiplayer**: Play against random opponents in real-time using WebRTC DataChannel.
- **WebAssembly**: The game logic is implemented in Go and compiled to WebAssembly for high performance.
- **Interactive UI**: Smooth and interactive user interface built with Ebiten.

## Getting Started

### Prerequisites

Deploy the following services as backend services:
- https://github.com/ponyo877/easy-matchmaking
  - Matchmaking server to connect users
- https://github.com/ponyo877/easy-rating
  - Rating management server by ero-rating(using github.com/kortemy/elo-go)
- https://github.com/OpenAyame/ayame
  - Signaling server for WebRTC connections


### Installation

1. Clone the repository & Install dependencies:

    ```sh
    git clone https://github.com/ponyo877/ebiten-hit-and-blow.git
    cd ebiten-hit-and-blow
    go mod tidy
    ```

2. Create Makefile:

    ```Makefile
    LCLLDFLAGS := -X 'github.com/ponyo877/ebiten-hit-and-blow/conn.wsScheme=ws' \
                  -X 'github.com/ponyo877/ebiten-hit-and-blow/conn.httpScheme=http' \
                  -X 'github.com/ponyo877/ebiten-hit-and-blow/conn.matchmakingOrigin=localhost:8000' \
                  -X 'github.com/ponyo877/ebiten-hit-and-blow/conn.signalingOrigin=localhost:3000' \
                  -X 'github.com/ponyo877/ebiten-hit-and-blow/conn.ratingOrigin=localhost:8001'
    COMFLAGS  :=  -X 'github.com/ponyo877/ebiten-hit-and-blow/go-ayame.turnHost=sample.com:3478' \
                  -X 'github.com/ponyo877/ebiten-hit-and-blow/go-ayame.turnUser=xxxxx' \
                  -X 'github.com/ponyo877/ebiten-hit-and-blow/go-ayame.turnPass=xxxxx' \
                  -X 'github.com/ponyo877/ebiten-hit-and-blow/conn.solt=xxxxx'
    wasm: main.go
        GOOS=js GOARCH=wasm go build -trimpath -ldflags "-s -w $(LCLLDFLAGS) $(COMFLAGS)" -o go/demo.wasm $<
        gzip -9 go/demo.wasm
    ```

3. Build the Go WebAssembly binary:

    ```sh
    make wasm
    ```

4. Serve the application:
    You can use any static file server to serve the index.html file. For example, using http-server:
    ```sh
    # if you use npx
    npx http-server .
    ```

5. Open the application:
    Open your browser and navigate to http://localhost:8080 (or the port your server is running on).

## Usage

- **Start a Game**: Open the application and click "Start Game" to find a match.
- **Make a Guess**: Enter your guess using the on-screen keyboard and submit.
- **Win the Game**: The first player to correctly guess the opponent's hand wins.


## Acknowledgements

- [Ebitengine](https://ebiten.org/) - A dead simple 2D game library for Go.
- [Pion WebRTC](https://github.com/pion/webrtc) - Pure Go implementation of the WebRTC API.

## License

The source code used in `go-ayame` directory are licensed under the following licenses:
```
Copyright 2020 Kazuyuki Honda (hakobera)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Modifications copyright (C) 2024 ponyo877
```

## Contact

For any inquiries or feedback, please reach out to [@ponyo877](https://twitter.com/ponyo877) on Twitter.

Enjoy the game and happy coding!
