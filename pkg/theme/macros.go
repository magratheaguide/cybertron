package theme

var defaultMacros = []Item{
	{Name: `A_LOCKED_B`, Text: `<div class='button-large closed'>Closed</div></item>`},
	{Name: `A_MOVED_B`, Text: `<div class='button-large closed'>Moved</div></item>`},
	{Name: `A_POLLONLY_B`, Text: `<div class='button-large closed'>Closed</div></item>`},
	{Name: `A_POST`, Text: `<div class='button-large'>New Topic</div></item>`},
	{Name: `A_REPLY`, Text: `<div class='button-large'>Add Reply</div></item>`},
	{Name: `A_POLL`, Text: `<div class='button-large'>New Poll</div></item>`},
	{Name: `A_STAR`, Text: `<img src='//files.jcink.net/themes/default/red-pip.png' /></item>`},
	{Name: `B_HOT`, Text: `<img alt="No New Posts" title="No New Posts" src="//files.jcink.net/themes/default/red-cube-small.png" /></item>`},
	{Name: `B_HOT_NN`, Text: `<img alt="No New Posts" title="No New Posts" src="//files.jcink.net/themes/default/white-cube-small.png" /></item>`},
	{Name: `B_LOCKED`, Text: `<img alt="Closed" title="Closed" src="//files.jcink.net/themes/default/closed.png" /></item>`},
	{Name: `B_MOVED`, Text: `<img alt="Moved Topic" title="Moved Topic" src='//files.jcink.net/themes/default/red-moved.png' /></item>`},
	{Name: `B_NEW`, Text: `<img alt="New Posts" title="New Posts" src="//files.jcink.net/themes/default/red-cube-small.png" /></item>`},
	{Name: `B_NORM`, Text: `<img alt="No New Posts" title="No New Posts" src="//files.jcink.net/themes/default/white-cube-small.png" /></item>`},
	{Name: `B_PIN`, Text: `<img alt="Pinned" title="Pinned" src="//files.jcink.net/themes/default/white-cube-small.png" /></item>`},
	{Name: `B_POLL`, Text: `<img alt="New Votes / Posts" title="New Votes / Posts" src="//files.jcink.net/themes/default/red-cube-small.png" /></item>`},
	{Name: `B_POLL_NN`, Text: `<img alt="No New Votes / Posts" title="No New Votes / Posts" src="//files.jcink.net/themes/default/white-cube-small.png" /></item>`},
	{Name: `B_HOT_DOT`, Text: `<img alt="New Posts" title="New Posts" src="//files.jcink.net/themes/default/red-cube-small.png" /></item>`},
	{Name: `B_NEW_DOT`, Text: `<img alt="New Posts" src="//files.jcink.net/themes/default/red-cube-small.png"></item>`},
	{Name: `B_HOT_NN_DOT`, Text: `<img alt="New Posts" title="New Posts"  src="//files.jcink.net/themes/default/white-cube-small.png"></item>`},
	{Name: `B_NORM_DOT`, Text: `<img alt="No New Posts" title="No New Posts"  src="//files.jcink.net/themes/default/white-cube-small.png"></item>`},
	{Name: `B_POLL_DOT`, Text: `<img alt="New Posts" title="New Posts" src="//files.jcink.net/themes/default/red-cube-small.png"></item>`},
	{Name: `B_POLL_NN_DOT`, Text: `<img alt="No New Posts" alt="No New Posts" src="//files.jcink.net/themes/default/white-cube-small.png"></item>`},
	{Name: `C_LOCKED`, Text: `<img alt="Read Only" title="Read Only" src="//files.jcink.net/themes/default/red-read-only.png" /></item>`},
	{Name: `C_OFF`, Text: `<img alt="No New Posts" title="No New Posts" src="//files.jcink.net/themes/default/white-cube-large.png" /></item>`},
	{Name: `C_OFF_CAT`, Text: `<img alt="No New Posts" title="No New Posts" src="//files.jcink.net/themes/default/white-cube-large.png" /></item>`},
	{Name: `C_OFF_RES`, Text: `<img alt="No New Posts" title="No New Posts" src="//files.jcink.net/themes/default/white-cube-large.png" /></item>`},
	{Name: `C_ON`, Text: `<img alt="New Posts" title="New Posts" src="//files.jcink.net/themes/default/red-cube-large.png" /></item>`},
	{Name: `C_ON_CAT`, Text: `<img alt="New Posts" title="New Posts" src="//files.jcink.net/themes/default/red-cube-large.png" /></item>`},
	{Name: `C_ON_RES`, Text: `<img alt="New Posts" title="New Posts" src="//files.jcink.net/themes/default/red-cube-large.png" /></item>`},
	{Name: `F_ACTIVE`, Text: `<img alt="Users" title="Users" src="//files.jcink.net/themes/default/active-users.png" /></item>`},
	{Name: `F_NAV_SEP`, Text: `&nbsp;&rarr;&nbsp;</item>`},
	{Name: `F_NAV`, Text: ` &raquo; </item>`},
	{Name: `F_STATS`, Text: `<img alt="Statistics" title="Statistics" src="//files.jcink.net/themes/default/statistics.png" /></item>`},
	{Name: `NO_PHOTO`, Text: `<img src='//files.jcink.net/themes/default/active-users.png'></item>`},
	{Name: `CAMERA`, Text: `View</item>`},
	{Name: `M_READ`, Text: `<img alt="Read Message" title="Read Message" src="//files.jcink.net/themes/default/white-cube-small.png" /></item>`},
	{Name: `M_UNREAD`, Text: `<img alt="Unread Message" title="Unread Message" src="//files.jcink.net/themes/default/red-cube-small.png" /></item>`},
	{Name: `P_AOL`, Text: `<span class='button-small'>AIM</span></item>`},
	{Name: `P_DELETE`, Text: `<span class='button-small'>Delete</span></item>`},
	{Name: `P_EDIT`, Text: `<span class='button-small'>Edit</span></item>`},
	{Name: `P_EMAIL`, Text: `<span class='button-small'>Email</span></item>`},
	{Name: `P_ICQ`, Text: `<span class='button-small'>GTalk</span></item>`},
	{Name: `P_MSG`, Text: `<span class='button-small'>PM</span></item>`},
	{Name: `P_QUOTE`, Text: `<span class='button-small'>Quote</span></item>`},
	{Name: `P_WEBSITE`, Text: `<span class='button-small'>Website</span></item>`},
	{Name: `P_YIM`, Text: `<span class='button-small'>YIM</span></item>`},
	{Name: `P_REPORT`, Text: `<span class='button-small'>Report</span></item>`},
	{Name: `P_MSN`, Text: `<span class='button-small'>MSN</span></item>`},
	{Name: `CAT_IMG`, Text: `</item>`},
	{Name: `NEW_POST`, Text: `<img src="//files.jcink.net/themes/default/double-arrow.gif" /></item>`},
	{Name: `LAST_POST`, Text: `<img src="//files.jcink.net/themes/default/double-arrow.gif" /></item>`},
	{Name: `BR_REDIRECT`, Text: `<img alt="Redirect" title="Redirect" src="//files.jcink.net/themes/default/red-redirect.png" /></item>`},
	{Name: `INTEGRITY_MSGR`, Text: `<span class='button-small'>Skype</span></item>`},
	{Name: `ADDRESS_CARD`, Text: `<span class='button-small'>Card</span></item>`},
	{Name: `T_QREPLY`, Text: `<div class='button-large'>Fast Reply</div></item>`},
	{Name: `T_OPTS`, Text: `<div class='button-large'>Topic Options</div></item>`},
	{Name: `CAL_NEWEVENT`, Text: `<div class='button-large'>Add Event</div></item>`},
	{Name: `F_RULES`, Text: `<img src='//files.jcink.net/html/acp/rules.png' alt='Forum Rules'></item>`},
	{Name: `WARN_0`, Text: ` <img src='//files.jcink.net/themes/default/gray-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> </item>`},
	{Name: `WARN_1`, Text: ` <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> </item>`},
	{Name: `WARN_2`, Text: ` <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> </item>`},
	{Name: `WARN_3`, Text: ` <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> </item>`},
	{Name: `WARN_4`, Text: ` <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/gray-block.png'> </item>`},
	{Name: `WARN_5`, Text: ` <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/red-block.png'> <img src='//files.jcink.net/themes/default/red-block.png'> </item>`},
	{Name: `WARN_ADD`, Text: `<img src='//files.jcink.net/themes/default/red-block-plus.png'></item>`},
	{Name: `WARN_MINUS`, Text: `<img src='//files.jcink.net/themes/default/red-block-minus.png'></item>`},
	{Name: `LOGO`, Text: `<img width='207' height='57' src="style_images/1/spacer.gif" /><!-- You can insert a logo image here in place of the white spacer --></item>`},
	{Name: `CUST_PROF`, Text: `<div class="postbit"><div align='center'><!-- |avatar| --><br><!-- |title| --><br></div></div><div class="postbit"><img src='//files.jcink.net/themes/default/users.png'> Group: <!-- |member_group| --> </div><div class="postbit"><img src='//files.jcink.net/themes/default/posts.png'> Posts: <!-- |posts| --></div><div class="postbit"><img src='//files.jcink.net/themes/default/joined.png'> Joined: <!-- |member_joined| --></div><div class="postbit">Age: <!-- |age| --><br>Location: <!-- |location| --><br>Status: <!-- |status| --><br></div><br><div style="margin-left: 14px;"><!-- |warn_text| --> <!-- |warn_minus| --><!-- |warn_img| --><!-- |warn_add| --></div></item>`},
	{Name: `B_PLUS_SHOW`, Text: `<img alt='Expand' title='Expand' src='//files.jcink.net/themes/default/cat-expand.png' /></item>`},
	{Name: `B_MINUS_HIDE`, Text: `<img alt='Collapse' title='Collapse' src='//files.jcink.net/themes/default/cat-collapse.png' /></item>`},
}

var allowedMacros = []string{
	"ADDRESS_CARD",
	"INTEGRITY_MSGR",
	"BR_REDIRECT",
	"LAST_POST",
	"NEW_POST",
	"CAT_IMG",
	"P_MSN",
	"P_REPORT",
	"P_YIM",
	"P_WEBSITE",
	"P_QUOTE",
	"P_MSG",
	"P_ICQ",
	"P_EMAIL",
	"P_EDIT",
	"P_DELETE",
	"P_AOL",
	"M_UNREAD",
	"M_READ",
	"CAMERA",
	"NO_PHOTO",
	"F_STATS",
	"F_NAV_SEP",
	"F_NAV",
	"F_ACTIVE",
	"C_ON_RES",
	"C_ON_CAT",
	"C_ON",
	"C_OFF_RES",
	"C_OFF_CAT",
	"C_OFF",
	"C_LOCKED",
	"B_POLL_NN_DOT",
	"B_POLL_DOT",
	"B_NORM_DOT",
	"B_HOT_NN_DOT",
	"B_NEW_DOT",
	"B_HOT_DOT",
	"B_POLL",
	"B_POLL_NN",
	"B_PIN",
	"B_NORM",
	"B_NEW",
	"B_MOVED",
	"B_LOCKED",
	"B_HOT_NN",
	"B_HOT",
	"A_STAR",
	"A_POLL",
	"A_REPLY",
	"A_POST",
	"A_POLLONLY_B",
	"A_MOVED_B",
	"A_LOCKED_B",
	"T_QREPLY",
	"T_OPTS",
	"CAL_NEWEVENT",
	"F_RULES",
	"WARN_0",
	"WARN_1",
	"WARN_2",
	"WARN_3",
	"WARN_4",
	"WARN_5",
	"WARN_ADD",
	"WARN_MINUS",
	"LOGO",
	"CUST_PROF",
	"P_ICON1",
	"P_ICON2",
	"P_ICON3",
	"P_ICON4",
	"P_ICON5",
	"P_ICON6",
	"P_ICON7",
	"P_ICON8",
	"P_ICON9",
	"P_ICON10",
	"P_ICON11",
	"P_ICON12",
	"P_ICON13",
	"P_ICON14",
	"T_TOPIC_ANNOUNCEMENT",
	"P_UP",
	"SIG_SEPARATOR",
	"T_ANNO_SUBHEAD",
	"T_PIN_SUBHEAD",
	"T_REG_SUBHEAD",
	"B_PLUS_SHOW",
	"B_MINUS_HIDE",
	"REP_PLUS",
	"REP_MINUS",
	"SUBLINKS_NEW",
	"SUBLINKS_NONEW",
	"SUB_SEPCHAR",
	"ALERT_NEW",
	"AVATAR",
	"AVATAR_URL",
}
