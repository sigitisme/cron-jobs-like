# cron-jobs-like
Using beego toolbox task to initiate cron jobs like functions.

by **sigitisme**

 - perMinuteTask -> runs every minutes but starts only if the second is 00
 - perSecondTask -> runs every seconds and starts immediately

since it only works if the app runs continously, func main() must use looping. for {} and beego.Run() works both.  
