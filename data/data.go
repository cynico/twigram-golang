package data

const (
	CodeVerifier    = "XDEADS"
	RedirectURI     = "http://timemusthaveastop.ddns.net/callback"
	ClientID        = "cXZxVklITjZwR2s4a1dDbW9uRDQ6MTpjaQ"
	ClientBasicAuth = "Basic Y1haeFZrbElUalp3UjJzNGExZERiVzl1UkRRNk1UcGphUTp2SEsyLWE4d0pNN1FJdWJPTm9UUFNfME50YzQtMVhWM042dTB0YVI0SjNaNkRkMWIxMA=="
	ClientSecret    = "vHK2-a8wJM7QIubONoTPS_0Ntc4-1XV3N6u0taR4J3Z6Dd1b10"
	SuccessPath     = "./public/views/success.html"
	FailurePath     = "./public/views/failure.html"
)

const (
	HelpMessage           = "/follow - Choose users by their twitter usernames to follow.\n\n/unfollow - Unfollow users you've previously followed.\n\n/list - Lits all the users you are currently following.\n\n/whereis - Check on a user you've not seen for a while by getting their last tweet.\n\n/unfollowall - Unfollow all the people you're currently following.\n\n/help - List all the commands and their jobs."
	FollowMessage         = "Send the username of the person you want to follow."
	UnfollowMessage       = "Send the username of the person you want to unfollow."
	UnrecognizableCommand = "There's no such command."
	TelegramToken         = "1871169654:AAG76d-JXlCWQ2pq8Y5lsmz93RzVFAX_ik8"
)
