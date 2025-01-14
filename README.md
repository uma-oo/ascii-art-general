Ascii-Art-Web-Export
====================

Overview
--------
Ascii-Art-Web-Export is a web application that allows users to generate ASCII art via a web interface and export the result as a text file (.txt format). The application ensures that users can download the generated ASCII art using proper HTTP headers to securely and efficiently transfer the file.

This project is built with Go, following best practices for HTTP handling, file export, and error management.


## Authors:
  **ooumayma**\
  **aayoubst**
  

Features
--------
- ASCII Art Export: Users can generate ASCII art and export the result in a .txt file.
- Download Button: A button or link on the webpage allows users to download the generated ASCII art.
- Proper HTTP Headers: Ensures file transfer using HTTP headers such as `Content-Type`, `Content-Length`, and `Content-Disposition`.
- Error Handling: Handles errors effectively to ensure a smooth user experience.
- Permissions: Files are exported with correct read and write permissions for the user.

Getting Started
---------------
### Prerequisites
To run this project, you will need Go installed on your system. You can download Go from https://golang.org/dl/.

### Installation
1. Clone the repository:
   git clone https://github.com/uma-oo/ascii-art-export-file
   cd ascii-art-export-file

2. Install dependencies:
   - This project only uses Go's standard library, so no external dependencies are required.

### Running the Web Application
To run the web server, execute the following command:

   go run ./app
The server will start at `http://localhost:6500`, where you can generate and export ASCII art.

### Exporting ASCII Art
- After generating the ASCII art on the webpage, you will see a button or link that allows you to download the art as a .txt file.
- The download is triggered by an HTTP endpoint that sets the correct headers to handle the file transfer.

HTTP Headers for File Export
-----------------------------
When exporting the ASCII art, the following HTTP headers will be used:

- Content-Type: text/plain
  - Specifies that the content is plain text.
  
- Content-Length: [file size]
  - Indicates the size of the file being transferred.

- Content-Disposition: attachment; filename="ascii-art.txt"
  - Suggests that the file should be downloaded with the name "ascii-art.txt".

Error Handling
--------------
The application is designed to handle errors gracefully. If there is an issue generating the ASCII art or exporting the file, the user will receive a relevant error message. For server errors, a 500 HTTP status code is returned along with an appropriate message.

Code Structure
--------------
The project follows a simple structure for easy understanding and maintenance:

ascii-art-web-export-file\
├── app             
| └── handler\
|          
|
├── assets          
│   └── index.css \
└── README.txt     

- main.go: This file contains the server logic, handling ASCII art generation and file export.
- index.html: This HTML file contains the web interface for generating ASCII art and exporting it.

Good Practices Followed
------------------------
- Modular Code: Functions are modular and separated logically.
- Error Handling: Errors are captured and communicated to the user.
- HTTP Headers: Proper HTTP headers are used for file export and content management.
- Code Comments: The code is well-commented to explain key sections.

Troubleshooting
---------------
- If you encounter issues with generating or exporting ASCII art, check the server logs for error details.
- Ensure the web server is running on http://localhost:6500
- If the download link or button doesn't work, verify that the export route is correctly set up in the Go code.



Contributing
------------
If you encounter issues or have suggestions for improvements, feel free to open an issue or submit a pull request. Contributions are welcome!

Acknowledgments
---------------
- Thanks to Go (Golang) for providing a simple platform for web development.
- Inspired by ASCII art enthusiasts.
- ChatGPT for generating this good README File
