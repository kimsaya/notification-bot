
DB => fileBase
1. Main DB (For overview all series, compiling DB_3 as row and DB_2 as column ) DB_1
2. User DB (Trace User Work) DB_2
3. Series DB (Trace Series Work) DB_3

Command List
//User Command
1. !set @series [Days] => In case worker choosing what they want to do.
1. => !set @series [Days] by@someone => Set someone to do.
Setting series time interval to [days]
2. !translate @series [Chapter X] @editor
-Adding work done by user (To DB_2)
-Update @series Chapter X translated by user (to DB_3)
- ignore the @editor (just to remind the another person)
3. !edit @series [Chapter X] @translator
-Same as ^
4. !post @series [Chapter X] 
-Same as ^
//Manager Command
5. !check @series [translate/edit/post]
-pull details of @series from DB3
7. !status @user [X Days ago]
-pull details of @user from DB_2
8. !assign @user [@series]
-assign @user to @series
9. !delay @series [X days]
-delay @series deadline by X day
10. !scan 
-a backup command to manually do function 1
-in case bot lost count of date upon reboot and will run comparison between due date and last post date to check if due date is not 0. If 0, will alert Function 1.

Function
1. Check if user last input for @series is more than due interval, then ping user
Bot: @translator, your @series Chapter X is due.