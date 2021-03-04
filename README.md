<p align="center">
  <img src="https://avatars1.githubusercontent.com/u/57274804?s=400&u=058242df13e206950c08efd68a540445ce4da17f&v=4" width="100">
</p>

# Confetti Commands

Use this template for your CLI project. Particularly suitable for a CI project.

> This code contains a few files to get started on your CLI project. If you want more functionalities, it is better to use [Confetti Framework](https://github.com/confetti-framework/confetti).

## Documentation

Confetti Commands has extensive and thorough [documentation](https://www.confetti-framework.com/docs/digging-deeper/commands), making it a breeze to get started.

## Example

``` go
package commands

import (
	"src/app/support"
	"github.com/confetti-framework/contract/inter"
	"io"
)

type SendEmails struct {
	Email string `flag:"email" required:"true"`
}

func (s SendEmails) Name() string {
	return "mail:send"
}

func (s SendEmails) Description() string {
	return "Send a marketing email to a user."
}

func (s SendEmails) Handle(c inter.Cli) inter.ExitCode {
	mailer := support.DripEmailer{}
	mailer.Send(s.Email)

	c.Info("Email send to: %s", s.Email)

	return inter.Success
}
```

## License

Confetti is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
