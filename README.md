# EchoGen

### what is this? (v0.1)

user will describe/give prompt for any topic they want to learn about and select option from how they want response from option like podcast, youtuber type, tutorial, etc. and will submit the request. and will get response in audio format, that user can listen to or download the audio file.

### TODO (v0.1)

- [ ] frontend
    - [x] simple input form(input, options, btn)
    - [x] send request to backend
    - [x] recive audio file and show it to user
    - [ ] show audio in place of sample prompt
    - [ ] google oauth
    
- [ ] backend
    - [x] receive request from frontend
    - [x] generate text content from prompt (use gemini api/openrouter)
    - [x] generate audio file from text content (elevenlabs api)
    - [x] send audio file to frontend
    - [ ] connect db
    - [ ] add auth