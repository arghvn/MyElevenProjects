creating a bot for calcualting age by using GO

at first we create an account in api.salck.com 
and in app section we create new app from scratch by name age-bot in Myrl Tech workspace 
in socket mode section we create socke-token instead token name
bot is created !!!

we can in add bot user event add another properties to our bot like history and message.im (for reading messages)and 
messages.channel and messages.mpin ...
if we use a lot of this properties , an error reduce in our app

another event chat:write (send messages) , chat:write:public (send messages to channles) , channels:read (view basic info about public channels in a workspace) ,im:read (view basic info about direct messages that age-bot has been added to) ,im:write (start direct messages with people) ,mpim:read (view basic info about group direct messages that age-bot has been added to) , mpim:write (start group direct messages with people)

so install the workspace and the bot installed in my worksap
now we have bot user oath token to access this workspace

after write bot code in main.go file we go to slack channel and in terminal(IDE) we write go run main.go for connecting to slack with socket mode
after connecting in slack channel we write age bot 
and write my yob is 1999 
oh it says your age is 23 correctly