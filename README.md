# EchoGen

### what is this? (v0.1)

EchoGen lets you learn about anything in a way that feels natural to you. Just type in a topic, pick how you want it explained—like a podcast, YouTuber, or step-by-step tutorial—and EchoGen will turn it into an audio you can listen to or download.

### TODO (v0.1)

- [ ] frontend
    - [x] simple input form(input, options, btn)
    - [x] send request to backend
    - [x] recive audio file and show it to user
    - [x] google oauth
    - [x] audio history
    
- [ ] backend
    - [x] receive request from frontend
    - [x] generate text content from prompt (use gemini api/openrouter)
    - [x] generate audio file from text content (elevenlabs api)
    - [x] send audio file to frontend
    - [x] connect db
    - [x] save values in db
        - [x] user
        - [x] audio generation
        - [x] save audio file to blob storage
    - [x] add auth
    - [x] get audio history from db
    - [x] add middleware
    - [x] add postgres
    - [x] add vercel blob storage