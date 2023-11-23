# Use targeting via bitmap index


### Challenge: Getting campaigns with different target sets

There are campaign sets, each with a list of allowed and/or blocked browsers, countries, devices and operating systems.

We have query criteria for the selector with browser, country, device, OS values. The selector should only return campaigns that match the query values based on allowed and blocked lists for each campaign.

#### For example: 

Campaign: ID: 11, BrowserAllowedList: [Chrome, IE], CountryBlockedList: [RU, EN], DeviceAllowedList[Mobile, TV], OSAllowedList[Android, IOS].

Campaign with ID 11 should be matched on request values: Browser: Chrome, Country: ES, Device: TV, OS: Android

And should not be matched on request values:  Browser: Chrome, Country: RU, Device: TV, OS: Android
 
