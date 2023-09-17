# go-zeptomail

`go-zeptomail` is a Go library for sending emails using the [ZeptoMail](https://zeptomail.com) API with HTML templates.

## Installation

You can install the package using `go get`:

```bash
go get github.com/blessedmadukoma/go-zeptomail
```

## Setup

Rename the `.env.example` to `.env` and replace the dummy data with your [ZeptoMail](https://zeptomail.com) credentails

## Usage:
```go
package main

import (
    "fmt"
    zeptomail "github.com/blessedmadukoma/go-zeptomail"
)

func main() {
  // smtp uses the ZeptoMail configurations provided in your dashboard. Store in .env file
  smtp := zeptomail.SMTP{
    Host: "smtp.zeptomail.com"
    Port: 465
    Username: "zeptomailusername"
    Password: "zeptomailpassword"
    SenderEmail: "zeptomail@mail.com"
  }

  data = zeptomail.MailData{
    RecipientName:  "Your recepient name",
    RecipientEmail: "email@mail.com",
    TemplateFile:   "welcome.html",
  }
    
  client := zeptomail.New(smtp)

  err := client.Send(data)
  if err != nil {
      fmt.Println(err)
      return
  }
}

```

## Contributing

If you'd like to contribute to this project, please follow these guidelines:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and test them thoroughly.
4. Commit your changes with clear commit messages.
5. Push your changes to your fork.
6. Create a pull request to the main repository.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

1. Thanks to ZeptoMail for providing the email-sending service.
2. This project is maintained by [Blessed Madukoma](github.com/blessedmadukoma).