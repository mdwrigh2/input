package input

const (
	KEY_RESERVED EventCode = 0
	KEY_ESC              = 1
	KEY_1                = 2
	KEY_2                = 3
	KEY_3                = 4
	KEY_4                = 5
	KEY_5                = 6
	KEY_6                = 7
	KEY_7                = 8
	KEY_8                = 9
	KEY_9                = 10
	KEY_0                = 11
	KEY_MINUS            = 12
	KEY_EQUAL            = 13
	KEY_BACKSPACE        = 14
	KEY_TAB              = 15
	KEY_Q                = 16
	KEY_W                = 17
	KEY_E                = 18
	KEY_R                = 19
	KEY_T                = 20
	KEY_Y                = 21
	KEY_U                = 22
	KEY_I                = 23
	KEY_O                = 24
	KEY_P                = 25
	KEY_LEFTBRACE        = 26
	KEY_RIGHTBRACE       = 27
	KEY_ENTER            = 28
	KEY_LEFTCTRL         = 29
	KEY_A                = 30
	KEY_S                = 31
	KEY_D                = 32
	KEY_F                = 33
	KEY_G                = 34
	KEY_H                = 35
	KEY_J                = 36
	KEY_K                = 37
	KEY_L                = 38
	KEY_SEMICOLON        = 39
	KEY_APOSTROPHE       = 40
	KEY_GRAVE            = 41
	KEY_LEFTSHIFT        = 42
	KEY_BACKSLASH        = 43
	KEY_Z                = 44
	KEY_X                = 45
	KEY_C                = 46
	KEY_V                = 47
	KEY_B                = 48
	KEY_N                = 49
	KEY_M                = 50
	KEY_COMMA            = 51
	KEY_DOT              = 52
	KEY_SLASH            = 53
	KEY_RIGHTSHIFT       = 54
	KEY_KPASTERISK       = 55
	KEY_LEFTALT          = 56
	KEY_SPACE            = 57
	KEY_CAPSLOCK         = 58
	KEY_F1               = 59
	KEY_F2               = 60
	KEY_F3               = 61
	KEY_F4               = 62
	KEY_F5               = 63
	KEY_F6               = 64
	KEY_F7               = 65
	KEY_F8               = 66
	KEY_F9               = 67
	KEY_F10              = 68
	KEY_NUMLOCK          = 69
	KEY_SCROLLLOCK       = 70
	KEY_KP7              = 71
	KEY_KP8              = 72
	KEY_KP9              = 73
	KEY_KPMINUS          = 74
	KEY_KP4              = 75
	KEY_KP5              = 76
	KEY_KP6              = 77
	KEY_KPPLUS           = 78
	KEY_KP1              = 79
	KEY_KP2              = 80
	KEY_KP3              = 81
	KEY_KP0              = 82
	KEY_KPDOT            = 83
	KEY_ZENKAKUHANKAKU   = 85
	KEY_102ND            = 86
	KEY_F11              = 87
	KEY_F12              = 88
	KEY_RO               = 89
	KEY_KATAKANA         = 90
	KEY_HIRAGANA         = 91
	KEY_HENKAN           = 92
	KEY_KATAKANAHIRAGANA = 93
	KEY_MUHENKAN         = 94
	KEY_KPJPCOMMA        = 95
	KEY_KPENTER          = 96
	KEY_RIGHTCTRL        = 97
	KEY_KPSLASH          = 98
	KEY_SYSRQ            = 99
	KEY_RIGHTALT         = 100
	KEY_LINEFEED         = 101
	KEY_HOME             = 102
	KEY_UP               = 103
	KEY_PAGEUP           = 104
	KEY_LEFT             = 105
	KEY_RIGHT            = 106
	KEY_END              = 107
	KEY_DOWN             = 108
	KEY_PAGEDOWN         = 109
	KEY_INSERT           = 110
	KEY_DELETE           = 111
	KEY_MACRO            = 112
	KEY_MUTE             = 113
	KEY_VOLUMEDOWN       = 114
	KEY_VOLUMEUP         = 115
	KEY_POWER            = 116 /* SC System Power Down */
	KEY_KPEQUAL          = 117
	KEY_KPPLUSMINUS      = 118
	KEY_PAUSE            = 119
	KEY_SCALE            = 120 /* AL Compiz Scale (Expose) */
	KEY_KPCOMMA          = 121
	KEY_HANGEUL          = 122
	KEY_HANGUEL          = KEY_HANGEUL
	KEY_HANJA            = 123
	KEY_YEN              = 124
	KEY_LEFTMETA         = 125
	KEY_RIGHTMETA        = 126
	KEY_COMPOSE          = 127
	KEY_STOP             = 128 /* AC Stop */
	KEY_AGAIN            = 129
	KEY_PROPS            = 130 /* AC Properties */
	KEY_UNDO             = 131 /* AC Undo */
	KEY_FRONT            = 132
	KEY_COPY             = 133 /* AC Copy */
	KEY_OPEN             = 134 /* AC Open */
	KEY_PASTE            = 135 /* AC Paste */
	KEY_FIND             = 136 /* AC Search */
	KEY_CUT              = 137 /* AC Cut */
	KEY_HELP             = 138 /* AL Integrated Help Center */
	KEY_MENU             = 139 /* Menu (show menu) */
	KEY_CALC             = 140 /* AL Calculator */
	KEY_SETUP            = 141
	KEY_SLEEP            = 142 /* SC System Sleep */
	KEY_WAKEUP           = 143 /* System Wake Up */
	KEY_FILE             = 144 /* AL Local Machine Browser */
	KEY_SENDFILE         = 145
	KEY_DELETEFILE       = 146
	KEY_XFER             = 147
	KEY_PROG1            = 148
	KEY_PROG2            = 149
	KEY_WWW              = 150 /* AL Internet Browser */
	KEY_MSDOS            = 151
	KEY_COFFEE           = 152 /* AL Terminal Lock/Screensaver */
	KEY_SCREENLOCK       = KEY_COFFEE
	KEY_DIRECTION        = 153
	KEY_CYCLEWINDOWS     = 154
	KEY_MAIL             = 155
	KEY_BOOKMARKS        = 156 /* AC Bookmarks */
	KEY_COMPUTER         = 157
	KEY_BACK             = 158 /* AC Back */
	KEY_FORWARD          = 159 /* AC Forward */
	KEY_CLOSECD          = 160
	KEY_EJECTCD          = 161
	KEY_EJECTCLOSECD     = 162
	KEY_NEXTSONG         = 163
	KEY_PLAYPAUSE        = 164
	KEY_PREVIOUSSONG     = 165
	KEY_STOPCD           = 166
	KEY_RECORD           = 167
	KEY_REWIND           = 168
	KEY_PHONE            = 169 /* Media Select Telephone */
	KEY_ISO              = 170
	KEY_CONFIG           = 171 /* AL Consumer Control Configuration */
	KEY_HOMEPAGE         = 172 /* AC Home */
	KEY_REFRESH          = 173 /* AC Refresh */
	KEY_EXIT             = 174 /* AC Exit */
	KEY_MOVE             = 175
	KEY_EDIT             = 176
	KEY_SCROLLUP         = 177
	KEY_SCROLLDOWN       = 178
	KEY_KPLEFTPAREN      = 179
	KEY_KPRIGHTPAREN     = 180
	KEY_NEW              = 181 /* AC New */
	KEY_REDO             = 182 /* AC Redo/Repeat */
	KEY_F13              = 183
	KEY_F14              = 184
	KEY_F15              = 185
	KEY_F16              = 186
	KEY_F17              = 187
	KEY_F18              = 188
	KEY_F19              = 189
	KEY_F20              = 190
	KEY_F21              = 191
	KEY_F22              = 192
	KEY_F23              = 193
	KEY_F24              = 194
	KEY_PLAYCD           = 200
	KEY_PAUSECD          = 201
	KEY_PROG3            = 202
	KEY_PROG4            = 203
	KEY_DASHBOARD        = 204 /* AL Dashboard */
	KEY_SUSPEND          = 205
	KEY_CLOSE            = 206 /* AC Close */
	KEY_PLAY             = 207
	KEY_FASTFORWARD      = 208
	KEY_BASSBOOST        = 209
	KEY_PRINT            = 210 /* AC Print */
	KEY_HP               = 211
	KEY_CAMERA           = 212
	KEY_SOUND            = 213
	KEY_QUESTION         = 214
	KEY_EMAIL            = 215
	KEY_CHAT             = 216
	KEY_SEARCH           = 217
	KEY_CONNECT          = 218
	KEY_FINANCE          = 219 /* AL Checkbook/Finance */
	KEY_SPORT            = 220
	KEY_SHOP             = 221
	KEY_ALTERASE         = 222
	KEY_CANCEL           = 223 /* AC Cancel */
	KEY_BRIGHTNESSDOWN   = 224
	KEY_BRIGHTNESSUP     = 225
	KEY_MEDIA            = 226
	KEY_SWITCHVIDEOMODE  = 227 /* Cycle between available video
	outputs (Monitor/LCD/TV-out/etc) */
	KEY_KBDILLUMTOGGLE   = 228
	KEY_KBDILLUMDOWN     = 229
	KEY_KBDILLUMUP       = 230
	KEY_SEND             = 231 /* AC Send */
	KEY_REPLY            = 232 /* AC Reply */
	KEY_FORWARDMAIL      = 233 /* AC Forward Msg */
	KEY_SAVE             = 234 /* AC Save */
	KEY_DOCUMENTS        = 235
	KEY_BATTERY          = 236
	KEY_BLUETOOTH        = 237
	KEY_WLAN             = 238
	KEY_UWB              = 239
	KEY_UNKNOWN          = 240
	KEY_VIDEO_NEXT       = 241 /* drive next video source */
	KEY_VIDEO_PREV       = 242 /* drive previous video source */
	KEY_BRIGHTNESS_CYCLE = 243 /* brightness up, after max is min */
	KEY_BRIGHTNESS_ZERO  = 244 /* brightness off, use ambient */
	KEY_DISPLAY_OFF      = 245 /* display device to off state */
	KEY_WIMAX            = 246
	KEY_RFKILL           = 247 /* Key that controls all radios */
	KEY_MICMUTE          = 248 /* Mute / unmute the microphone */
	BTN_MISC             = 0x100
	BTN_0                = 0x100
	BTN_1                = 0x101
	BTN_2                = 0x102
	BTN_3                = 0x103
	BTN_4                = 0x104
	BTN_5                = 0x105
	BTN_6                = 0x106
	BTN_7                = 0x107
	BTN_8                = 0x108
	BTN_9                = 0x109
	BTN_MOUSE            = 0x110
	BTN_LEFT             = 0x110
	BTN_RIGHT            = 0x111
	BTN_MIDDLE           = 0x112
	BTN_SIDE             = 0x113
	BTN_EXTRA            = 0x114
	BTN_FORWARD          = 0x115
	BTN_BACK             = 0x116
	BTN_TASK             = 0x117
	BTN_JOYSTICK         = 0x120
	BTN_TRIGGER          = 0x120
	BTN_THUMB            = 0x121
	BTN_THUMB2           = 0x122
	BTN_TOP              = 0x123
	BTN_TOP2             = 0x124
	BTN_PINKIE           = 0x125
	BTN_BASE             = 0x126
	BTN_BASE2            = 0x127
	BTN_BASE3            = 0x128
	BTN_BASE4            = 0x129
	BTN_BASE5            = 0x12a
	BTN_BASE6            = 0x12b
	BTN_DEAD             = 0x12f
	BTN_GAMEPAD          = 0x130
	BTN_A                = 0x130
	BTN_B                = 0x131
	BTN_C                = 0x132
	BTN_X                = 0x133
	BTN_Y                = 0x134
	BTN_Z                = 0x135
	BTN_TL               = 0x136
	BTN_TR               = 0x137
	BTN_TL2              = 0x138
	BTN_TR2              = 0x139
	BTN_SELECT           = 0x13a
	BTN_START            = 0x13b
	BTN_MODE             = 0x13c
	BTN_THUMBL           = 0x13d
	BTN_THUMBR           = 0x13e
	BTN_DIGI             = 0x140
	BTN_TOOL_PEN         = 0x140
	BTN_TOOL_RUBBER      = 0x141
	BTN_TOOL_BRUSH       = 0x142
	BTN_TOOL_PENCIL      = 0x143
	BTN_TOOL_AIRBRUSH    = 0x144
	BTN_TOOL_FINGER      = 0x145
	BTN_TOOL_MOUSE       = 0x146
	BTN_TOOL_LENS        = 0x147
	BTN_TOOL_QUINTTAP    = 0x148 /* Five fingers on trackpad */
	BTN_TOUCH            = 0x14a
	BTN_STYLUS           = 0x14b
	BTN_STYLUS2          = 0x14c
	BTN_TOOL_DOUBLETAP   = 0x14d
	BTN_TOOL_TRIPLETAP   = 0x14e
	BTN_TOOL_QUADTAP     = 0x14f /* Four fingers on trackpad */
	BTN_WHEEL            = 0x150
	BTN_GEAR_DOWN        = 0x150
	BTN_GEAR_UP          = 0x151
	KEY_OK               = 0x160
	KEY_SELECT           = 0x161
	KEY_GOTO             = 0x162
	KEY_CLEAR            = 0x163
	KEY_POWER2           = 0x164
	KEY_OPTION           = 0x165
	KEY_INFO             = 0x166 /* AL OEM Features/Tips/Tutorial */
	KEY_TIME             = 0x167
	KEY_VENDOR           = 0x168
	KEY_ARCHIVE          = 0x169
	KEY_PROGRAM          = 0x16a /* Media Select Program Guide */
	KEY_CHANNEL          = 0x16b
	KEY_FAVORITES        = 0x16c
	KEY_EPG              = 0x16d
	KEY_PVR              = 0x16e /* Media Select Home */
	KEY_MHP              = 0x16f
	KEY_LANGUAGE         = 0x170
	KEY_TITLE            = 0x171
	KEY_SUBTITLE         = 0x172
	KEY_ANGLE            = 0x173
	KEY_ZOOM             = 0x174
	KEY_MODE             = 0x175
	KEY_KEYBOARD         = 0x176
	KEY_SCREEN           = 0x177
	KEY_PC               = 0x178 /* Media Select Computer */
	KEY_TV               = 0x179 /* Media Select TV */
	KEY_TV2              = 0x17a /* Media Select Cable */
	KEY_VCR              = 0x17b /* Media Select VCR */
	KEY_VCR2             = 0x17c /* VCR Plus */
	KEY_SAT              = 0x17d /* Media Select Satellite */
	KEY_SAT2             = 0x17e
	KEY_CD               = 0x17f /* Media Select CD */
	KEY_TAPE             = 0x180 /* Media Select Tape */
	KEY_RADIO            = 0x181
	KEY_TUNER            = 0x182 /* Media Select Tuner */
	KEY_PLAYER           = 0x183
	KEY_TEXT             = 0x184
	KEY_DVD              = 0x185 /* Media Select DVD */
	KEY_AUX              = 0x186
	KEY_MP3              = 0x187
	KEY_AUDIO            = 0x188 /* AL Audio Browser */
	KEY_VIDEO            = 0x189 /* AL Movie Browser */
	KEY_DIRECTORY        = 0x18a
	KEY_LIST             = 0x18b
	KEY_MEMO             = 0x18c /* Media Select Messages */
	KEY_CALENDAR         = 0x18d
	KEY_RED              = 0x18e
	KEY_GREEN            = 0x18f
	KEY_YELLOW           = 0x190
	KEY_BLUE             = 0x191
	KEY_CHANNELUP        = 0x192 /* Channel Increment */
	KEY_CHANNELDOWN      = 0x193 /* Channel Decrement */
	KEY_FIRST            = 0x194
	KEY_LAST             = 0x195 /* Recall Last */
	KEY_AB               = 0x196
	KEY_NEXT             = 0x197
	KEY_RESTART          = 0x198
	KEY_SLOW             = 0x199
	KEY_SHUFFLE          = 0x19a
	KEY_BREAK            = 0x19b
	KEY_PREVIOUS         = 0x19c
	KEY_DIGITS           = 0x19d
	KEY_TEEN             = 0x19e
	KEY_TWEN             = 0x19f
	KEY_VIDEOPHONE       = 0x1a0 /* Media Select Video Phone */
	KEY_GAMES            = 0x1a1 /* Media Select Games */
	KEY_ZOOMIN           = 0x1a2 /* AC Zoom In */
	KEY_ZOOMOUT          = 0x1a3 /* AC Zoom Out */
	KEY_ZOOMRESET        = 0x1a4 /* AC Zoom */
	KEY_WORDPROCESSOR    = 0x1a5 /* AL Word Processor */
	KEY_EDITOR           = 0x1a6 /* AL Text Editor */
	KEY_SPREADSHEET      = 0x1a7 /* AL Spreadsheet */
	KEY_GRAPHICSEDITOR   = 0x1a8 /* AL Graphics Editor */
	KEY_PRESENTATION     = 0x1a9 /* AL Presentation App */
	KEY_DATABASE         = 0x1aa /* AL Database App */
	KEY_NEWS             = 0x1ab /* AL Newsreader */
	KEY_VOICEMAIL        = 0x1ac /* AL Voicemail */
	KEY_ADDRESSBOOK      = 0x1ad /* AL Contacts/Address Book */
	KEY_MESSENGER        = 0x1ae /* AL Instant Messaging */
	KEY_DISPLAYTOGGLE    = 0x1af /* Turn display (LCD) on and off */
	KEY_SPELLCHECK       = 0x1b0 /* AL Spell Check */
	KEY_LOGOFF           = 0x1b1 /* AL Logoff */
	KEY_DOLLAR           = 0x1b2
	KEY_EURO             = 0x1b3
	KEY_FRAMEBACK        = 0x1b4 /* Consumer - transport controls */
	KEY_FRAMEFORWARD     = 0x1b5
	KEY_CONTEXT_MENU     = 0x1b6 /* GenDesc - system context menu */
	KEY_MEDIA_REPEAT     = 0x1b7 /* Consumer - transport control */
	KEY_10CHANNELSUP     = 0x1b8 /* 10 channels up (10+) */
	KEY_10CHANNELSDOWN   = 0x1b9 /* 10 channels down (10-) */
	KEY_IMAGES           = 0x1ba /* AL Image Browser */
	KEY_DEL_EOL          = 0x1c0
	KEY_DEL_EOS          = 0x1c1
	KEY_INS_LINE         = 0x1c2
	KEY_DEL_LINE         = 0x1c3
	KEY_FN               = 0x1d0
	KEY_FN_ESC           = 0x1d1
	KEY_FN_F1            = 0x1d2
	KEY_FN_F2            = 0x1d3
	KEY_FN_F3            = 0x1d4
	KEY_FN_F4            = 0x1d5
	KEY_FN_F5            = 0x1d6
	KEY_FN_F6            = 0x1d7
	KEY_FN_F7            = 0x1d8
	KEY_FN_F8            = 0x1d9
	KEY_FN_F9            = 0x1da
	KEY_FN_F10           = 0x1db
	KEY_FN_F11           = 0x1dc
	KEY_FN_F12           = 0x1dd
	KEY_FN_1             = 0x1de
	KEY_FN_2             = 0x1df
	KEY_FN_D             = 0x1e0
	KEY_FN_E             = 0x1e1
	KEY_FN_F             = 0x1e2
	KEY_FN_S             = 0x1e3
	KEY_FN_B             = 0x1e4
	KEY_BRL_DOT1         = 0x1f1
	KEY_BRL_DOT2         = 0x1f2
	KEY_BRL_DOT3         = 0x1f3
	KEY_BRL_DOT4         = 0x1f4
	KEY_BRL_DOT5         = 0x1f5
	KEY_BRL_DOT6         = 0x1f6
	KEY_BRL_DOT7         = 0x1f7
	KEY_BRL_DOT8         = 0x1f8
	KEY_BRL_DOT9         = 0x1f9
	KEY_BRL_DOT10        = 0x1fa
	KEY_NUMERIC_0        = 0x200 /* used by phones, remote controls, */
	KEY_NUMERIC_1        = 0x201 /* and other keypads */
	KEY_NUMERIC_2        = 0x202
	KEY_NUMERIC_3        = 0x203
	KEY_NUMERIC_4        = 0x204
	KEY_NUMERIC_5        = 0x205
	KEY_NUMERIC_6        = 0x206
	KEY_NUMERIC_7        = 0x207
	KEY_NUMERIC_8        = 0x208
	KEY_NUMERIC_9        = 0x209
	KEY_NUMERIC_STAR     = 0x20a
	KEY_NUMERIC_POUND    = 0x20b
	KEY_CAMERA_FOCUS     = 0x210
	KEY_WPS_BUTTON       = 0x211 /* WiFi Protected Setup key */
	KEY_TOUCHPAD_TOGGLE  = 0x212 /* Request switch touchpad on or off */
	KEY_TOUCHPAD_ON      = 0x213
	KEY_TOUCHPAD_OFF     = 0x214
	KEY_CAMERA_ZOOMIN    = 0x215
	KEY_CAMERA_ZOOMOUT   = 0x216
	KEY_CAMERA_UP        = 0x217
	KEY_CAMERA_DOWN      = 0x218
	KEY_CAMERA_LEFT      = 0x219
	KEY_CAMERA_RIGHT     = 0x21a
	BTN_TRIGGER_HAPPY    = 0x2c0
	BTN_TRIGGER_HAPPY1   = 0x2c0
	BTN_TRIGGER_HAPPY2   = 0x2c1
	BTN_TRIGGER_HAPPY3   = 0x2c2
	BTN_TRIGGER_HAPPY4   = 0x2c3
	BTN_TRIGGER_HAPPY5   = 0x2c4
	BTN_TRIGGER_HAPPY6   = 0x2c5
	BTN_TRIGGER_HAPPY7   = 0x2c6
	BTN_TRIGGER_HAPPY8   = 0x2c7
	BTN_TRIGGER_HAPPY9   = 0x2c8
	BTN_TRIGGER_HAPPY10  = 0x2c9
	BTN_TRIGGER_HAPPY11  = 0x2ca
	BTN_TRIGGER_HAPPY12  = 0x2cb
	BTN_TRIGGER_HAPPY13  = 0x2cc
	BTN_TRIGGER_HAPPY14  = 0x2cd
	BTN_TRIGGER_HAPPY15  = 0x2ce
	BTN_TRIGGER_HAPPY16  = 0x2cf
	BTN_TRIGGER_HAPPY17  = 0x2d0
	BTN_TRIGGER_HAPPY18  = 0x2d1
	BTN_TRIGGER_HAPPY19  = 0x2d2
	BTN_TRIGGER_HAPPY20  = 0x2d3
	BTN_TRIGGER_HAPPY21  = 0x2d4
	BTN_TRIGGER_HAPPY22  = 0x2d5
	BTN_TRIGGER_HAPPY23  = 0x2d6
	BTN_TRIGGER_HAPPY24  = 0x2d7
	BTN_TRIGGER_HAPPY25  = 0x2d8
	BTN_TRIGGER_HAPPY26  = 0x2d9
	BTN_TRIGGER_HAPPY27  = 0x2da
	BTN_TRIGGER_HAPPY28  = 0x2db
	BTN_TRIGGER_HAPPY29  = 0x2dc
	BTN_TRIGGER_HAPPY30  = 0x2dd
	BTN_TRIGGER_HAPPY31  = 0x2de
	BTN_TRIGGER_HAPPY32  = 0x2df
	BTN_TRIGGER_HAPPY33  = 0x2e0
	BTN_TRIGGER_HAPPY34  = 0x2e1
	BTN_TRIGGER_HAPPY35  = 0x2e2
	BTN_TRIGGER_HAPPY36  = 0x2e3
	BTN_TRIGGER_HAPPY37  = 0x2e4
	BTN_TRIGGER_HAPPY38  = 0x2e5
	BTN_TRIGGER_HAPPY39  = 0x2e6
	BTN_TRIGGER_HAPPY40  = 0x2e7
	KEY_MIN_INTERESTING  = KEY_MUTE
	KEY_MAX              = 0x2ff
	KEY_CNT              = (KEY_MAX + 1)
	REL_X                = 0x00
	REL_Y                = 0x01
	REL_Z                = 0x02
	REL_RX               = 0x03
	REL_RY               = 0x04
	REL_RZ               = 0x05
	REL_HWHEEL           = 0x06
	REL_DIAL             = 0x07
	REL_WHEEL            = 0x08
	REL_MISC             = 0x09
	REL_MAX              = 0x0f
	REL_CNT              = (REL_MAX + 1)
	ABS_X                = 0x00
	ABS_Y                = 0x01
	ABS_Z                = 0x02
	ABS_RX               = 0x03
	ABS_RY               = 0x04
	ABS_RZ               = 0x05
	ABS_THROTTLE         = 0x06
	ABS_RUDDER           = 0x07
	ABS_WHEEL            = 0x08
	ABS_GAS              = 0x09
	ABS_BRAKE            = 0x0a
	ABS_HAT0X            = 0x10
	ABS_HAT0Y            = 0x11
	ABS_HAT1X            = 0x12
	ABS_HAT1Y            = 0x13
	ABS_HAT2X            = 0x14
	ABS_HAT2Y            = 0x15
	ABS_HAT3X            = 0x16
	ABS_HAT3Y            = 0x17
	ABS_PRESSURE         = 0x18
	ABS_DISTANCE         = 0x19
	ABS_TILT_X           = 0x1a
	ABS_TILT_Y           = 0x1b
	ABS_TOOL_WIDTH       = 0x1c
	ABS_VOLUME           = 0x20
	ABS_MISC             = 0x28
	ABS_MT_SLOT          = 0x2f /* MT slot being modified */
	ABS_MT_TOUCH_MAJOR   = 0x30 /* Major axis of touching ellipse */
	ABS_MT_TOUCH_MINOR   = 0x31 /* Minor axis (omit if circular) */
	ABS_MT_WIDTH_MAJOR   = 0x32 /* Major axis of approaching ellipse */
	ABS_MT_WIDTH_MINOR   = 0x33 /* Minor axis (omit if circular) */
	ABS_MT_ORIENTATION   = 0x34 /* Ellipse orientation */
	ABS_MT_POSITION_X    = 0x35 /* Center X ellipse position */
	ABS_MT_POSITION_Y    = 0x36 /* Center Y ellipse position */
	ABS_MT_TOOL_TYPE     = 0x37 /* Type of touching device */
	ABS_MT_BLOB_ID       = 0x38 /* Group a set of packets as a blob */
	ABS_MT_TRACKING_ID   = 0x39 /* Unique ID of initiated contact */
	ABS_MT_PRESSURE      = 0x3a /* Pressure on contact area */
	ABS_MT_DISTANCE      = 0x3b /* Contact hover distance */
	ABS_MAX              = 0x3f
	ABS_CNT              = (ABS_MAX + 1)
	SW_LID               = 0x00 /* set = lid shut */
	SW_TABLET_MODE       = 0x01 /* set = tablet mode */
	SW_HEADPHONE_INSERT  = 0x02 /* set = inserted */
	SW_RFKILL_ALL        = 0x03 /* rfkill master switch, type "any"
	set = radio enabled */
	SW_RADIO                = SW_RFKILL_ALL /* deprecated */
	SW_MICROPHONE_INSERT    = 0x04          /* set = inserted */
	SW_DOCK                 = 0x05          /* set = plugged into dock */
	SW_LINEOUT_INSERT       = 0x06          /* set = inserted */
	SW_JACK_PHYSICAL_INSERT = 0x07          /* set = mechanical switch set */
	SW_VIDEOOUT_INSERT      = 0x08          /* set = inserted */
	SW_CAMERA_LENS_COVER    = 0x09          /* set = lens covered */
	SW_KEYPAD_SLIDE         = 0x0a          /* set = keypad slide out */
	SW_FRONT_PROXIMITY      = 0x0b          /* set = front proximity sensor active */
	SW_ROTATE_LOCK          = 0x0c          /* set = rotate locked/disabled */
	SW_LINEIN_INSERT        = 0x0d          /* set = inserted */
	SW_MAX                  = 0x0f
	SW_CNT                  = (SW_MAX + 1)
	MSC_SERIAL              = 0x00
	MSC_PULSELED            = 0x01
	MSC_GESTURE             = 0x02
	MSC_RAW                 = 0x03
	MSC_SCAN                = 0x04
	MSC_MAX                 = 0x07
	MSC_CNT                 = (MSC_MAX + 1)
	LED_NUML                = 0x00
	LED_CAPSL               = 0x01
	LED_SCROLLL             = 0x02
	LED_COMPOSE             = 0x03
	LED_KANA                = 0x04
	LED_SLEEP               = 0x05
	LED_SUSPEND             = 0x06
	LED_MUTE                = 0x07
	LED_MISC                = 0x08
	LED_MAIL                = 0x09
	LED_CHARGING            = 0x0a
	LED_MAX                 = 0x0f
	LED_CNT                 = (LED_MAX + 1)
	REP_DELAY               = 0x00
	REP_PERIOD              = 0x01
	REP_MAX                 = 0x01
	REP_CNT                 = (REP_MAX + 1)
	SND_CLICK               = 0x00
	SND_BELL                = 0x01
	SND_TONE                = 0x02
	SND_MAX                 = 0x07
	SND_CNT                 = (SND_MAX + 1)
	ID_BUS                  = 0
	ID_VENDOR               = 1
	ID_PRODUCT              = 2
	ID_VERSION              = 3
	BUS_PCI                 = 0x01
	BUS_ISAPNP              = 0x02
	BUS_USB                 = 0x03
	BUS_HIL                 = 0x04
	BUS_BLUETOOTH           = 0x05
	BUS_VIRTUAL             = 0x06
	BUS_ISA                 = 0x10
	BUS_I8042               = 0x11
	BUS_XTKBD               = 0x12
	BUS_RS232               = 0x13
	BUS_GAMEPORT            = 0x14
	BUS_PARPORT             = 0x15
	BUS_AMIGA               = 0x16
	BUS_ADB                 = 0x17
	BUS_I2C                 = 0x18
	BUS_HOST                = 0x19
	BUS_GSC                 = 0x1A
	BUS_ATARI               = 0x1B
	BUS_SPI                 = 0x1C
	MT_TOOL_FINGER          = 0
	MT_TOOL_PEN             = 1
	MT_TOOL_MAX             = 1
	FF_STATUS_STOPPED       = 0x00
	FF_STATUS_PLAYING       = 0x01
	FF_STATUS_MAX           = 0x01
)

var (
	KeycodeToName = map[EventCode]string{
		0:             "KEY_RESERVED",
		1:             "KEY_ESC",
		2:             "KEY_1",
		3:             "KEY_2",
		4:             "KEY_3",
		5:             "KEY_4",
		6:             "KEY_5",
		7:             "KEY_6",
		8:             "KEY_7",
		9:             "KEY_8",
		10:            "KEY_9",
		11:            "KEY_0",
		12:            "KEY_MINUS",
		13:            "KEY_EQUAL",
		14:            "KEY_BACKSPACE",
		15:            "KEY_TAB",
		16:            "KEY_Q",
		17:            "KEY_W",
		18:            "KEY_E",
		19:            "KEY_R",
		20:            "KEY_T",
		21:            "KEY_Y",
		22:            "KEY_U",
		23:            "KEY_I",
		24:            "KEY_O",
		25:            "KEY_P",
		26:            "KEY_LEFTBRACE",
		27:            "KEY_RIGHTBRACE",
		28:            "KEY_ENTER",
		29:            "KEY_LEFTCTRL",
		30:            "KEY_A",
		31:            "KEY_S",
		32:            "KEY_D",
		33:            "KEY_F",
		34:            "KEY_G",
		35:            "KEY_H",
		36:            "KEY_J",
		37:            "KEY_K",
		38:            "KEY_L",
		39:            "KEY_SEMICOLON",
		40:            "KEY_APOSTROPHE",
		41:            "KEY_GRAVE",
		42:            "KEY_LEFTSHIFT",
		43:            "KEY_BACKSLASH",
		44:            "KEY_Z",
		45:            "KEY_X",
		46:            "KEY_C",
		47:            "KEY_V",
		48:            "KEY_B",
		49:            "KEY_N",
		50:            "KEY_M",
		51:            "KEY_COMMA",
		52:            "KEY_DOT",
		53:            "KEY_SLASH",
		54:            "KEY_RIGHTSHIFT",
		55:            "KEY_KPASTERISK",
		56:            "KEY_LEFTALT",
		57:            "KEY_SPACE",
		58:            "KEY_CAPSLOCK",
		59:            "KEY_F1",
		60:            "KEY_F2",
		61:            "KEY_F3",
		62:            "KEY_F4",
		63:            "KEY_F5",
		64:            "KEY_F6",
		65:            "KEY_F7",
		66:            "KEY_F8",
		67:            "KEY_F9",
		68:            "KEY_F10",
		69:            "KEY_NUMLOCK",
		70:            "KEY_SCROLLLOCK",
		71:            "KEY_KP7",
		72:            "KEY_KP8",
		73:            "KEY_KP9",
		74:            "KEY_KPMINUS",
		75:            "KEY_KP4",
		76:            "KEY_KP5",
		77:            "KEY_KP6",
		78:            "KEY_KPPLUS",
		79:            "KEY_KP1",
		80:            "KEY_KP2",
		81:            "KEY_KP3",
		82:            "KEY_KP0",
		83:            "KEY_KPDOT",
		85:            "KEY_ZENKAKUHANKAKU",
		86:            "KEY_102ND",
		87:            "KEY_F11",
		88:            "KEY_F12",
		89:            "KEY_RO",
		90:            "KEY_KATAKANA",
		91:            "KEY_HIRAGANA",
		92:            "KEY_HENKAN",
		93:            "KEY_KATAKANAHIRAGANA",
		94:            "KEY_MUHENKAN",
		95:            "KEY_KPJPCOMMA",
		96:            "KEY_KPENTER",
		97:            "KEY_RIGHTCTRL",
		98:            "KEY_KPSLASH",
		99:            "KEY_SYSRQ",
		100:           "KEY_RIGHTALT",
		101:           "KEY_LINEFEED",
		102:           "KEY_HOME",
		103:           "KEY_UP",
		104:           "KEY_PAGEUP",
		105:           "KEY_LEFT",
		106:           "KEY_RIGHT",
		107:           "KEY_END",
		108:           "KEY_DOWN",
		109:           "KEY_PAGEDOWN",
		110:           "KEY_INSERT",
		111:           "KEY_DELETE",
		112:           "KEY_MACRO",
		113:           "KEY_MUTE",
		114:           "KEY_VOLUMEDOWN",
		115:           "KEY_VOLUMEUP",
		116:           "KEY_POWER", /* SC System Power Down */
		117:           "KEY_KPEQUAL",
		118:           "KEY_KPPLUSMINUS",
		119:           "KEY_PAUSE",
		120:           "KEY_SCALE", /* AL Compiz Scale (Expose) */
		121:           "KEY_KPCOMMA",
		122:           "KEY_HANGEUL",
		123:           "KEY_HANJA",
		124:           "KEY_YEN",
		125:           "KEY_LEFTMETA",
		126:           "KEY_RIGHTMETA",
		127:           "KEY_COMPOSE",
		128:           "KEY_STOP", /* AC Stop */
		129:           "KEY_AGAIN",
		130:           "KEY_PROPS", /* AC Properties */
		131:           "KEY_UNDO",  /* AC Undo */
		132:           "KEY_FRONT",
		133:           "KEY_COPY",  /* AC Copy */
		134:           "KEY_OPEN",  /* AC Open */
		135:           "KEY_PASTE", /* AC Paste */
		136:           "KEY_FIND",  /* AC Search */
		137:           "KEY_CUT",   /* AC Cut */
		138:           "KEY_HELP",  /* AL Integrated Help Center */
		139:           "KEY_MENU",  /* Menu (show menu) */
		140:           "KEY_CALC",  /* AL Calculator */
		141:           "KEY_SETUP",
		142:           "KEY_SLEEP",  /* SC System Sleep */
		143:           "KEY_WAKEUP", /* System Wake Up */
		144:           "KEY_FILE",   /* AL Local Machine Browser */
		145:           "KEY_SENDFILE",
		146:           "KEY_DELETEFILE",
		147:           "KEY_XFER",
		148:           "KEY_PROG1",
		149:           "KEY_PROG2",
		150:           "KEY_WWW", /* AL Internet Browser */
		151:           "KEY_MSDOS",
		152:           "KEY_COFFEE", /* AL Terminal Lock/Screensaver */
		153:           "KEY_DIRECTION",
		154:           "KEY_CYCLEWINDOWS",
		155:           "KEY_MAIL",
		156:           "KEY_BOOKMARKS", /* AC Bookmarks */
		157:           "KEY_COMPUTER",
		158:           "KEY_BACK",    /* AC Back */
		159:           "KEY_FORWARD", /* AC Forward */
		160:           "KEY_CLOSECD",
		161:           "KEY_EJECTCD",
		162:           "KEY_EJECTCLOSECD",
		163:           "KEY_NEXTSONG",
		164:           "KEY_PLAYPAUSE",
		165:           "KEY_PREVIOUSSONG",
		166:           "KEY_STOPCD",
		167:           "KEY_RECORD",
		168:           "KEY_REWIND",
		169:           "KEY_PHONE", /* Media Select Telephone */
		170:           "KEY_ISO",
		171:           "KEY_CONFIG",   /* AL Consumer Control Configuration */
		172:           "KEY_HOMEPAGE", /* AC Home */
		173:           "KEY_REFRESH",  /* AC Refresh */
		174:           "KEY_EXIT",     /* AC Exit */
		175:           "KEY_MOVE",
		176:           "KEY_EDIT",
		177:           "KEY_SCROLLUP",
		178:           "KEY_SCROLLDOWN",
		179:           "KEY_KPLEFTPAREN",
		180:           "KEY_KPRIGHTPAREN",
		181:           "KEY_NEW",  /* AC New */
		182:           "KEY_REDO", /* AC Redo/Repeat */
		183:           "KEY_F13",
		184:           "KEY_F14",
		185:           "KEY_F15",
		186:           "KEY_F16",
		187:           "KEY_F17",
		188:           "KEY_F18",
		189:           "KEY_F19",
		190:           "KEY_F20",
		191:           "KEY_F21",
		192:           "KEY_F22",
		193:           "KEY_F23",
		194:           "KEY_F24",
		200:           "KEY_PLAYCD",
		201:           "KEY_PAUSECD",
		202:           "KEY_PROG3",
		203:           "KEY_PROG4",
		204:           "KEY_DASHBOARD", /* AL Dashboard */
		205:           "KEY_SUSPEND",
		206:           "KEY_CLOSE", /* AC Close */
		207:           "KEY_PLAY",
		208:           "KEY_FASTFORWARD",
		209:           "KEY_BASSBOOST",
		210:           "KEY_PRINT", /* AC Print */
		211:           "KEY_HP",
		212:           "KEY_CAMERA",
		213:           "KEY_SOUND",
		214:           "KEY_QUESTION",
		215:           "KEY_EMAIL",
		216:           "KEY_CHAT",
		217:           "KEY_SEARCH",
		218:           "KEY_CONNECT",
		219:           "KEY_FINANCE", /* AL Checkbook/Finance */
		220:           "KEY_SPORT",
		221:           "KEY_SHOP",
		222:           "KEY_ALTERASE",
		223:           "KEY_CANCEL", /* AC Cancel */
		224:           "KEY_BRIGHTNESSDOWN",
		225:           "KEY_BRIGHTNESSUP",
		226:           "KEY_MEDIA",
		227:           "KEY_SWITCHVIDEOMODE",
		228:           "KEY_KBDILLUMTOGGLE",
		229:           "KEY_KBDILLUMDOWN",
		230:           "KEY_KBDILLUMUP",
		231:           "KEY_SEND",        /* AC Send */
		232:           "KEY_REPLY",       /* AC Reply */
		233:           "KEY_FORWARDMAIL", /* AC Forward Msg */
		234:           "KEY_SAVE",        /* AC Save */
		235:           "KEY_DOCUMENTS",
		236:           "KEY_BATTERY",
		237:           "KEY_BLUETOOTH",
		238:           "KEY_WLAN",
		239:           "KEY_UWB",
		240:           "KEY_UNKNOWN",
		241:           "KEY_VIDEO_NEXT",       /* drive next video source */
		242:           "KEY_VIDEO_PREV",       /* drive previous video source */
		243:           "KEY_BRIGHTNESS_CYCLE", /* brightness up, after max is min */
		244:           "KEY_BRIGHTNESS_ZERO",  /* brightness off, use ambient */
		245:           "KEY_DISPLAY_OFF",      /* display device to off state */
		246:           "KEY_WIMAX",
		247:           "KEY_RFKILL",  /* Key that controls all radios */
		248:           "KEY_MICMUTE", /* Mute / unmute the microphone */
		0x100:         "BTN_0",
		0x101:         "BTN_1",
		0x102:         "BTN_2",
		0x103:         "BTN_3",
		0x104:         "BTN_4",
		0x105:         "BTN_5",
		0x106:         "BTN_6",
		0x107:         "BTN_7",
		0x108:         "BTN_8",
		0x109:         "BTN_9",
		0x110:         "BTN_MOUSE",
		0x111:         "BTN_RIGHT",
		0x112:         "BTN_MIDDLE",
		0x113:         "BTN_SIDE",
		0x114:         "BTN_EXTRA",
		0x115:         "BTN_FORWARD",
		0x116:         "BTN_BACK",
		0x117:         "BTN_TASK",
		0x120:         "BTN_JOYSTICK",
		0x121:         "BTN_THUMB",
		0x122:         "BTN_THUMB2",
		0x123:         "BTN_TOP",
		0x124:         "BTN_TOP2",
		0x125:         "BTN_PINKIE",
		0x126:         "BTN_BASE",
		0x127:         "BTN_BASE2",
		0x128:         "BTN_BASE3",
		0x129:         "BTN_BASE4",
		0x12a:         "BTN_BASE5",
		0x12b:         "BTN_BASE6",
		0x12f:         "BTN_DEAD",
		0x130:         "BTN_A",
		0x131:         "BTN_B",
		0x132:         "BTN_C",
		0x133:         "BTN_X",
		0x134:         "BTN_Y",
		0x135:         "BTN_Z",
		0x136:         "BTN_TL",
		0x137:         "BTN_TR",
		0x138:         "BTN_TL2",
		0x139:         "BTN_TR2",
		0x13a:         "BTN_SELECT",
		0x13b:         "BTN_START",
		0x13c:         "BTN_MODE",
		0x13d:         "BTN_THUMBL",
		0x13e:         "BTN_THUMBR",
		0x140:         "BTN_TOOL_PEN",
		0x141:         "BTN_TOOL_RUBBER",
		0x142:         "BTN_TOOL_BRUSH",
		0x143:         "BTN_TOOL_PENCIL",
		0x144:         "BTN_TOOL_AIRBRUSH",
		0x145:         "BTN_TOOL_FINGER",
		0x146:         "BTN_TOOL_MOUSE",
		0x147:         "BTN_TOOL_LENS",
		0x148:         "BTN_TOOL_QUINTTAP", /* Five fingers on trackpad */
		0x14a:         "BTN_TOUCH",
		0x14b:         "BTN_STYLUS",
		0x14c:         "BTN_STYLUS2",
		0x14d:         "BTN_TOOL_DOUBLETAP",
		0x14e:         "BTN_TOOL_TRIPLETAP",
		0x14f:         "BTN_TOOL_QUADTAP", /* Four fingers on trackpad */
		0x150:         "BTN_GEAR_DOWN",
		0x151:         "BTN_GEAR_UP",
		0x160:         "KEY_OK",
		0x161:         "KEY_SELECT",
		0x162:         "KEY_GOTO",
		0x163:         "KEY_CLEAR",
		0x164:         "KEY_POWER2",
		0x165:         "KEY_OPTION",
		0x166:         "KEY_INFO", /* AL OEM Features/Tips/Tutorial */
		0x167:         "KEY_TIME",
		0x168:         "KEY_VENDOR",
		0x169:         "KEY_ARCHIVE",
		0x16a:         "KEY_PROGRAM", /* Media Select Program Guide */
		0x16b:         "KEY_CHANNEL",
		0x16c:         "KEY_FAVORITES",
		0x16d:         "KEY_EPG",
		0x16e:         "KEY_PVR", /* Media Select Home */
		0x16f:         "KEY_MHP",
		0x170:         "KEY_LANGUAGE",
		0x171:         "KEY_TITLE",
		0x172:         "KEY_SUBTITLE",
		0x173:         "KEY_ANGLE",
		0x174:         "KEY_ZOOM",
		0x175:         "KEY_MODE",
		0x176:         "KEY_KEYBOARD",
		0x177:         "KEY_SCREEN",
		0x178:         "KEY_PC",   /* Media Select Computer */
		0x179:         "KEY_TV",   /* Media Select TV */
		0x17a:         "KEY_TV2",  /* Media Select Cable */
		0x17b:         "KEY_VCR",  /* Media Select VCR */
		0x17c:         "KEY_VCR2", /* VCR Plus */
		0x17d:         "KEY_SAT",  /* Media Select Satellite */
		0x17e:         "KEY_SAT2",
		0x17f:         "KEY_CD",   /* Media Select CD */
		0x180:         "KEY_TAPE", /* Media Select Tape */
		0x181:         "KEY_RADIO",
		0x182:         "KEY_TUNER", /* Media Select Tuner */
		0x183:         "KEY_PLAYER",
		0x184:         "KEY_TEXT",
		0x185:         "KEY_DVD", /* Media Select DVD */
		0x186:         "KEY_AUX",
		0x187:         "KEY_MP3",
		0x188:         "KEY_AUDIO", /* AL Audio Browser */
		0x189:         "KEY_VIDEO", /* AL Movie Browser */
		0x18a:         "KEY_DIRECTORY",
		0x18b:         "KEY_LIST",
		0x18c:         "KEY_MEMO", /* Media Select Messages */
		0x18d:         "KEY_CALENDAR",
		0x18e:         "KEY_RED",
		0x18f:         "KEY_GREEN",
		0x190:         "KEY_YELLOW",
		0x191:         "KEY_BLUE",
		0x192:         "KEY_CHANNELUP",   /* Channel Increment */
		0x193:         "KEY_CHANNELDOWN", /* Channel Decrement */
		0x194:         "KEY_FIRST",
		0x195:         "KEY_LAST", /* Recall Last */
		0x196:         "KEY_AB",
		0x197:         "KEY_NEXT",
		0x198:         "KEY_RESTART",
		0x199:         "KEY_SLOW",
		0x19a:         "KEY_SHUFFLE",
		0x19b:         "KEY_BREAK",
		0x19c:         "KEY_PREVIOUS",
		0x19d:         "KEY_DIGITS",
		0x19e:         "KEY_TEEN",
		0x19f:         "KEY_TWEN",
		0x1a0:         "KEY_VIDEOPHONE",     /* Media Select Video Phone */
		0x1a1:         "KEY_GAMES",          /* Media Select Games */
		0x1a2:         "KEY_ZOOMIN",         /* AC Zoom In */
		0x1a3:         "KEY_ZOOMOUT",        /* AC Zoom Out */
		0x1a4:         "KEY_ZOOMRESET",      /* AC Zoom */
		0x1a5:         "KEY_WORDPROCESSOR",  /* AL Word Processor */
		0x1a6:         "KEY_EDITOR",         /* AL Text Editor */
		0x1a7:         "KEY_SPREADSHEET",    /* AL Spreadsheet */
		0x1a8:         "KEY_GRAPHICSEDITOR", /* AL Graphics Editor */
		0x1a9:         "KEY_PRESENTATION",   /* AL Presentation App */
		0x1aa:         "KEY_DATABASE",       /* AL Database App */
		0x1ab:         "KEY_NEWS",           /* AL Newsreader */
		0x1ac:         "KEY_VOICEMAIL",      /* AL Voicemail */
		0x1ad:         "KEY_ADDRESSBOOK",    /* AL Contacts/Address Book */
		0x1ae:         "KEY_MESSENGER",      /* AL Instant Messaging */
		0x1af:         "KEY_DISPLAYTOGGLE",  /* Turn display (LCD) on and off */
		0x1b0:         "KEY_SPELLCHECK",     /* AL Spell Check */
		0x1b1:         "KEY_LOGOFF",         /* AL Logoff */
		0x1b2:         "KEY_DOLLAR",
		0x1b3:         "KEY_EURO",
		0x1b4:         "KEY_FRAMEBACK", /* Consumer - transport controls */
		0x1b5:         "KEY_FRAMEFORWARD",
		0x1b6:         "KEY_CONTEXT_MENU",   /* GenDesc - system context menu */
		0x1b7:         "KEY_MEDIA_REPEAT",   /* Consumer - transport control */
		0x1b8:         "KEY_10CHANNELSUP",   /* 10 channels up (10+) */
		0x1b9:         "KEY_10CHANNELSDOWN", /* 10 channels down (10-) */
		0x1ba:         "KEY_IMAGES",         /* AL Image Browser */
		0x1c0:         "KEY_DEL_EOL",
		0x1c1:         "KEY_DEL_EOS",
		0x1c2:         "KEY_INS_LINE",
		0x1c3:         "KEY_DEL_LINE",
		0x1d0:         "KEY_FN",
		0x1d1:         "KEY_FN_ESC",
		0x1d2:         "KEY_FN_F1",
		0x1d3:         "KEY_FN_F2",
		0x1d4:         "KEY_FN_F3",
		0x1d5:         "KEY_FN_F4",
		0x1d6:         "KEY_FN_F5",
		0x1d7:         "KEY_FN_F6",
		0x1d8:         "KEY_FN_F7",
		0x1d9:         "KEY_FN_F8",
		0x1da:         "KEY_FN_F9",
		0x1db:         "KEY_FN_F10",
		0x1dc:         "KEY_FN_F11",
		0x1dd:         "KEY_FN_F12",
		0x1de:         "KEY_FN_1",
		0x1df:         "KEY_FN_2",
		0x1e0:         "KEY_FN_D",
		0x1e1:         "KEY_FN_E",
		0x1e2:         "KEY_FN_F",
		0x1e3:         "KEY_FN_S",
		0x1e4:         "KEY_FN_B",
		0x1f1:         "KEY_BRL_DOT1",
		0x1f2:         "KEY_BRL_DOT2",
		0x1f3:         "KEY_BRL_DOT3",
		0x1f4:         "KEY_BRL_DOT4",
		0x1f5:         "KEY_BRL_DOT5",
		0x1f6:         "KEY_BRL_DOT6",
		0x1f7:         "KEY_BRL_DOT7",
		0x1f8:         "KEY_BRL_DOT8",
		0x1f9:         "KEY_BRL_DOT9",
		0x1fa:         "KEY_BRL_DOT10",
		0x200:         "KEY_NUMERIC_0", /* used by phones, remote controls, */
		0x201:         "KEY_NUMERIC_1", /* and other keypads */
		0x202:         "KEY_NUMERIC_2",
		0x203:         "KEY_NUMERIC_3",
		0x204:         "KEY_NUMERIC_4",
		0x205:         "KEY_NUMERIC_5",
		0x206:         "KEY_NUMERIC_6",
		0x207:         "KEY_NUMERIC_7",
		0x208:         "KEY_NUMERIC_8",
		0x209:         "KEY_NUMERIC_9",
		0x20a:         "KEY_NUMERIC_STAR",
		0x20b:         "KEY_NUMERIC_POUND",
		0x210:         "KEY_CAMERA_FOCUS",
		0x211:         "KEY_WPS_BUTTON",      /* WiFi Protected Setup key */
		0x212:         "KEY_TOUCHPAD_TOGGLE", /* Request switch touchpad on or off */
		0x213:         "KEY_TOUCHPAD_ON",
		0x214:         "KEY_TOUCHPAD_OFF",
		0x215:         "KEY_CAMERA_ZOOMIN",
		0x216:         "KEY_CAMERA_ZOOMOUT",
		0x217:         "KEY_CAMERA_UP",
		0x218:         "KEY_CAMERA_DOWN",
		0x219:         "KEY_CAMERA_LEFT",
		0x21a:         "KEY_CAMERA_RIGHT",
		0x2c0:         "BTN_TRIGGER_HAPPY1",
		0x2c1:         "BTN_TRIGGER_HAPPY2",
		0x2c2:         "BTN_TRIGGER_HAPPY3",
		0x2c3:         "BTN_TRIGGER_HAPPY4",
		0x2c4:         "BTN_TRIGGER_HAPPY5",
		0x2c5:         "BTN_TRIGGER_HAPPY6",
		0x2c6:         "BTN_TRIGGER_HAPPY7",
		0x2c7:         "BTN_TRIGGER_HAPPY8",
		0x2c8:         "BTN_TRIGGER_HAPPY9",
		0x2c9:         "BTN_TRIGGER_HAPPY10",
		0x2ca:         "BTN_TRIGGER_HAPPY11",
		0x2cb:         "BTN_TRIGGER_HAPPY12",
		0x2cc:         "BTN_TRIGGER_HAPPY13",
		0x2cd:         "BTN_TRIGGER_HAPPY14",
		0x2ce:         "BTN_TRIGGER_HAPPY15",
		0x2cf:         "BTN_TRIGGER_HAPPY16",
		0x2d0:         "BTN_TRIGGER_HAPPY17",
		0x2d1:         "BTN_TRIGGER_HAPPY18",
		0x2d2:         "BTN_TRIGGER_HAPPY19",
		0x2d3:         "BTN_TRIGGER_HAPPY20",
		0x2d4:         "BTN_TRIGGER_HAPPY21",
		0x2d5:         "BTN_TRIGGER_HAPPY22",
		0x2d6:         "BTN_TRIGGER_HAPPY23",
		0x2d7:         "BTN_TRIGGER_HAPPY24",
		0x2d8:         "BTN_TRIGGER_HAPPY25",
		0x2d9:         "BTN_TRIGGER_HAPPY26",
		0x2da:         "BTN_TRIGGER_HAPPY27",
		0x2db:         "BTN_TRIGGER_HAPPY28",
		0x2dc:         "BTN_TRIGGER_HAPPY29",
		0x2dd:         "BTN_TRIGGER_HAPPY30",
		0x2de:         "BTN_TRIGGER_HAPPY31",
		0x2df:         "BTN_TRIGGER_HAPPY32",
		0x2e0:         "BTN_TRIGGER_HAPPY33",
		0x2e1:         "BTN_TRIGGER_HAPPY34",
		0x2e2:         "BTN_TRIGGER_HAPPY35",
		0x2e3:         "BTN_TRIGGER_HAPPY36",
		0x2e4:         "BTN_TRIGGER_HAPPY37",
		0x2e5:         "BTN_TRIGGER_HAPPY38",
		0x2e6:         "BTN_TRIGGER_HAPPY39",
		0x2e7:         "BTN_TRIGGER_HAPPY40",
		0x2ff:         "KEY_MAX",
		(KEY_MAX + 1): "KEY_CNT",
	}

	RelCodeToName = map[EventCode]string{
		0x00: "REL_X",
		0x01: "REL_Y",
		0x02: "REL_Z",
		0x03: "REL_RX",
		0x04: "REL_RY",
		0x05: "REL_RZ",
		0x06: "REL_HWHEEL",
		0x07: "REL_DIAL",
		0x08: "REL_WHEEL",
		0x09: "REL_MISC",
		0x0f: "REL_MAX",
	}
	AbsCodeToName = map[EventCode]string{
		0x00: "ABS_X",
		0x01: "ABS_Y",
		0x02: "ABS_Z",
		0x03: "ABS_RX",
		0x04: "ABS_RY",
		0x05: "ABS_RZ",
		0x06: "ABS_THROTTLE",
		0x07: "ABS_RUDDER",
		0x08: "ABS_WHEEL",
		0x09: "ABS_GAS",
		0x0a: "ABS_BRAKE",
		0x10: "ABS_HAT0X",
		0x11: "ABS_HAT0Y",
		0x12: "ABS_HAT1X",
		0x13: "ABS_HAT1Y",
		0x14: "ABS_HAT2X",
		0x15: "ABS_HAT2Y",
		0x16: "ABS_HAT3X",
		0x17: "ABS_HAT3Y",
		0x18: "ABS_PRESSURE",
		0x19: "ABS_DISTANCE",
		0x1a: "ABS_TILT_X",
		0x1b: "ABS_TILT_Y",
		0x1c: "ABS_TOOL_WIDTH",
		0x20: "ABS_VOLUME",
		0x28: "ABS_MISC",
		0x2f: "ABS_MT_SLOT",        /* MT slot being modified */
		0x30: "ABS_MT_TOUCH_MAJOR", /* Major axis of touching ellipse */
		0x31: "ABS_MT_TOUCH_MINOR", /* Minor axis (omit if circular) */
		0x32: "ABS_MT_WIDTH_MAJOR", /* Major axis of approaching ellipse */
		0x33: "ABS_MT_WIDTH_MINOR", /* Minor axis (omit if circular) */
		0x34: "ABS_MT_ORIENTATION", /* Ellipse orientation */
		0x35: "ABS_MT_POSITION_X",  /* Center X ellipse position */
		0x36: "ABS_MT_POSITION_Y",  /* Center Y ellipse position */
		0x37: "ABS_MT_TOOL_TYPE",   /* Type of touching device */
		0x38: "ABS_MT_BLOB_ID",     /* Group a set of packets as a blob */
		0x39: "ABS_MT_TRACKING_ID", /* Unique ID of initiated contact */
		0x3a: "ABS_MT_PRESSURE",    /* Pressure on contact area */
		0x3b: "ABS_MT_DISTANCE",    /* Contact hover distance */
		0x3f: "ABS_MAX",
	}

	SwitchCodeToName = map[EventCode]string{
		0x00: "SW_LID",                  /* set = lid shut */
		0x01: "SW_TABLET_MODE",          /* set = tablet mode */
		0x02: "SW_HEADPHONE_INSERT",     /* set = inserted */
		0x03: "SW_RFKILL_ALL",           /* rfkill master switch, type any set radio enabled */
		0x04: "SW_MICROPHONE_INSERT",    /* set = inserted */
		0x05: "SW_DOCK",                 /* set = plugged into dock */
		0x06: "SW_LINEOUT_INSERT",       /* set = inserted */
		0x07: "SW_JACK_PHYSICAL_INSERT", /* set = mechanical switch set */
		0x08: "SW_VIDEOOUT_INSERT",      /* set = inserted */
		0x09: "SW_CAMERA_LENS_COVER",    /* set = lens covered */
		0x0a: "SW_KEYPAD_SLIDE",         /* set = keypad slide out */
		0x0b: "SW_FRONT_PROXIMITY",      /* set = front proximity sensor active */
		0x0c: "SW_ROTATE_LOCK",          /* set = rotate locked/disabled */
		0x0d: "SW_LINEIN_INSERT",        /* set = inserted */
		0x0f: "SW_MAX",
	}

	MscCodeToName = map[EventCode]string{
		0x00: "MSC_SERIAL",
		0x01: "MSC_PULSELED",
		0x02: "MSC_GESTURE",
		0x03: "MSC_RAW",
		0x04: "MSC_SCAN",
		0x07: "MSC_MAX",
	}

	LedCodeToName = map[EventCode]string{
		0x00: "LED_NUML",
		0x01: "LED_CAPSL",
		0x02: "LED_SCROLLL",
		0x03: "LED_COMPOSE",
		0x04: "LED_KANA",
		0x05: "LED_SLEEP",
		0x06: "LED_SUSPEND",
		0x07: "LED_MUTE",
		0x08: "LED_MISC",
		0x09: "LED_MAIL",
		0x0a: "LED_CHARGING",
		0x0f: "LED_MAX",
	}

	SoundCodeToName = map[EventCode]string{
		0x00: "SND_CLICK",
		0x01: "SND_BELL",
		0x02: "SND_TONE",
		0x07: "SND_MAX",
	}

	IdCodeToName = map[EventCode]string{
		0: "ID_BUS",
		1: "ID_VENDOR",
		2: "ID_PRODUCT",
		3: "ID_VERSION",
	}

	BusCodeToName = map[EventCode]string{
		0x01: "BUS_PCI",
		0x02: "BUS_ISAPNP",
		0x03: "BUS_USB",
		0x04: "BUS_HIL",
		0x05: "BUS_BLUETOOTH",
		0x06: "BUS_VIRTUAL",
		0x10: "BUS_ISA",
		0x11: "BUS_I8042",
		0x12: "BUS_XTKBD",
		0x13: "BUS_RS232",
		0x14: "BUS_GAMEPORT",
		0x15: "BUS_PARPORT",
		0x16: "BUS_AMIGA",
		0x17: "BUS_ADB",
		0x18: "BUS_I2C",
		0x19: "BUS_HOST",
		0x1A: "BUS_GSC",
		0x1B: "BUS_ATARI",
		0x1C: "BUS_SPI",
	}

	MultiTouchCodeToName = map[EventCode]string{
		0: "MT_TOOL_FINGER",
		1: "MT_TOOL_PEN",
	}

	ForceFeedbackCodeToName = map[EventCode]string{
		0x00: "FF_STATUS_STOPPED",
		0x01: "FF_STATUS_PLAYING",
	}
)
