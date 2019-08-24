# giveme

cli tool using go. to send emails and ask for my meals

## User stories

I want a command line app to order my meals. I have to send manualy an email each day to order them and I'd to doit in the cli without have to write so much:

- As a user I can choose between fixed menu or a custom order. not listed but available in the menu.
- As a user a have to enter my email, receptor email, name to send the email.
- As a user I want my last selections in memory and a command like send last order or something.

# tools
[cobra](https://github.com/spf13/cobra) - for the cli toolset
[promptui](https://github.com/manifoldco/promptui) - for selects and advanced prompt labels
smtp (native go package) - for emails
