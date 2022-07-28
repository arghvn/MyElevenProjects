building an email verifier tool 
checking that an email is true or fake

after making this code if we write go run main.go in terminal we have a statement
domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarRecord thath says 
when you enter a domain name right now it will give you the doamin name and check it and mx records...
and by checking these then we will know if that domain is right from which that email id is come out from to or not
so youll just check the domain
in the next part we enter an email record to check like mailchimp.com it returns a statement for us like 
mailchimp.com,true,false,,true,v=DMARC1; p = reject; rua = mailto:19ezfriw@ag.dmarcian.com; ruf = mailto:19ezfriw@fr.dmarcian.com
that means :
mailchimp.com is the domain , hasMX is true and etc(spfRecord is blank and no spfRecord)