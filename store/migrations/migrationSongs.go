package migrations

import "github.com/mizmorr/songslib/internal/model"

func returnSongs() []model.Song {
	songs := []model.Song{
		{
			Band:   "The Beatles",
			Name:   "Hey Jude",
			Lyrics: "Hey Jude, don't make it bad\nTake a sad song and make it better\nRemember to let her into your heart\nThen you can start to make it better",
		},
		{
			Band:   "The Beatles",
			Name:   "Help!",
			Lyrics: "Help me, please, help me now\nWith a loving heart\nI've been broken, but I'm trying to find my way\nBack to the light",
		},

		{
			Band: "Queen",
			Name: "Bohemian Rhapsody",
			Lyrics: `Is this the real life? Is this just fantasy?
Caught in a landslide, no escape from reality
Open your eyes, look up to the skies and see
I'm just a poor boy, I need no sympathy
Because I'm easy come, easy go
Little high, little low
Any way the wind blows doesn't really matter to me, to me`,
		},
		{
			Band: "Led Zeppelin",
			Name: "Stairway to Heaven",
			Lyrics: `There's a lady who's sure all that glitters is gold
And she's buying a stairway to Heaven
When she gets there she knows, if the stores are all closed
With a word she can get what she came for
Ooh, ooh, and she's buying a stairway to Heaven
There's a sign on the wall, but she wants to be sure`,
		},

		{
			Band:   "Nirvana",
			Name:   "Smells Like Teen Spirit",
			Lyrics: "With the lights out, it's less dangerous\nHere we are now, entertain us\n\nA mulatto, an albino, a mosquito, my libido\nYeah",
		},
		{
			Band:   "Radiohead",
			Name:   "Creep",
			Lyrics: "When you were here before\nCouldn't look you in the eye\n\nI'm a creep\nI'm a weirdo\nWhat the hell am I doing here?\nI don't belong here",
		},
		{
			Band:   "Pink Floyd",
			Name:   "Comfortably Numb",
			Lyrics: "Hello\nIs there anybody in there?\nJust nod if you can hear me\n\nThere is no pain you are receding\nA distant ship's smoke on the horizon",
		},
		{
			Band: "Ethereal Wanderers",
			Name: "Stairway of Dreams",
			Lyrics: `In the still of the night, a whisper takes flight
Guiding her steps to the edge of the light
A stairway of stars, glowing brighter each turn
With every step, her soul starts to burn`,
		},
		{
			Band: "Mystic Horizon",
			Name: "Echoes of the Sky",
			Lyrics: `Far beyond the hills, where the silence resides
A song in the distance, a beacon that guides
Through valleys of wonder, and forests untamed
She follows the echoes, her heart unclaimed`,
		},
		{
			Band: "Celestial Voyage",
			Name: "Path of Light",
			Lyrics: `In the depths of the night, where the shadows lay low
A path starts to shimmer, with a gentle glow
Each step is a story, each turn is a dream
She follows the light, where the heavens gleam`,
		},
		{
			Band:   "Black Sabbath",
			Name:   "Paranoid",
			Lyrics: "I'm waiting for the day\nWhen I'll be able to say\nThat I've been saved\nAnd there's no more pain.\nYou're gonna love me\nLike I love you\nOne and one is two\nOh, yeah.\nDon't you try to take me on\n'Cause you know I'm right\nIt's gonna take you too long\nTo realize I'm outta sight.\nSo don't bother me\nLeave me alone\nCan't you see\nI want to be\nOn my own.",
		},
		{
			Band:   "Black Sabbath",
			Name:   "Iron Man",
			Lyrics: "Stark indestructible\nObject of the scorned\nNever give an inch\nAlways got to have it your way.\nThey call you iron man\nHeavy metal thunder\nRunning through the night\nWith the speed of light.\nThere's a thousand faces\nHidden in the past\nEvery single one\nScreaming in the night.\nThe man behind the mask\nIs a master of disguise\nHe's the one who's got it\nAll figured out.",
		},
		{
			Band:   "Black Sabbath",
			Name:   "War Pigs",
			Lyrics: "Generals gathered in their masses\nJust like witches at black masses\nEvil minds that plot destruction\nSorcerers of death's construction.\nIn the fields the bodies burning\nAs the war machine keeps turning\nDeath and hatred to mankind\nPoisoning their brainwashed minds.",
		},
		{
			Band:   "Megadeth",
			Name:   "Symphony of Destruction",
			Lyrics: "I'm just a man, with a mission to serve\nAnd I've got a plan, it's a plan that will work\nI've been around, I've seen things you'd never believe\nBut I know my mind, and I know what I need\n\nSymphony of destruction\nPlaying in my head\nGot to get this message to you\nBefore I'm dead\n\nYou don't know me, you can't see me\nBut you feel my presence all the same\nYou can run but you can't hide\n'Cause I'm inside your brain\n\nSymphony of destruction\nPlaying in my head\nGot to get this message to you\nBefore I'm dead\n\nNow I'm breaking out, I won't be controlled\nThis is my life, I call the shots, I make the rules\nI won't take no shit, from anyone at all\nI do what I want, I say what I please, I am the law\n\nSymphony of destruction\nPlaying in my head\nGot to get this message to you\nBefore I'm dead",
		},
		{
			Band:   "Megadeth",
			Name:   "Peace Sells... But Who's Buying?",
			Lyrics: "Hey, hey, hey!\nWe're all living in black and white\nUntil we add a little color\nTo the grey side of our lives\nHey, hey, yeah!\nIt's time for you to realize\nThat things aren't always what they seem\nWhen you look beyond the lies\n\nPeace sells, but who's buying?\nNo one's there to help you pay\nThe price of peace is your soul\nSo be prepared to pay\n\nThey're coming to collect the debt\nOf all the things you never said\nAll the time you never spent\nWith the ones you left behind\n\nPeace sells, but who's buying?\nNo one's there to help you pay\nThe price of peace is your soul\nSo be prepared to pay\n\nDon't think you can hide away\nFrom the demons in your mind\nThey'll find you, they'll haunt you\nTill you're running out of time\n\nPeace sells, but who's buying?\nNo one's there to help you pay\nThe price of peace is your soul\nSo be prepared to pay",
		},
		{
			Band:   "Red Hot Chili Peppers",
			Name:   "Californication",
			Lyrics: "In the days of ancient Rome\nWe were outnumbered and afraid\nCaesar said, «I'm gonna make you feel safe»\nAs long as you give me all your money\nAnd now the workers are restless\nThey're tired of being oppressed\nBut I can assure you of this\nYou'll be fed, but you will not be impressed\nOh, I'm a real go-getter\nBaby, I've been on the streets\nI can take you to the top\nOr I can make you starve in the heat\nIt's called Californication\nLand of make believe\nIf you ask me how I'm livin'\nI could say, «Just a-wait and see»\nSo I don't wonder when I hear the news\nCalifornia's sufferin' from the usual blues\nThe earthquakes and the hurricanes\nHave got no place to run\nTelevangelists say\n«Come and get your healing done»\nOut in Hollywood they're always so cool\nThey got a lotta fancy cars\nBut they still got problems\nTryin' to find God\nIn the land of Californication\nWhere it seems that dreams come true\nThere's a man who's never seen snow\nHe's got a lot of money too\nPeople like to think\nThat they know what you're thinking\nWell, they don't have a clue\nCause if you ask them for some help\nThey might just tell you the truth\nIt's called Californication\nNothing's ever what it seems\nIt's just another illusion\nOn the silver screen\nNow everybody's acting like they're stars\nFame and fortune are just a state of mind\nEverybody's going to heaven\nBecause they got the right kind of faith\nSome people think that life is just a game\nOther people say that life's a mystery\nAll I know is that I've got to get away\nFrom this place of Californication\nBefore I lose my mind",
		},
		{
			Band:   "Red Hot Chili Peppers",
			Name:   "Under the Bridge",
			Lyrics: "I was born under the bridge\nMy mother was a crackhead\nNo one cared if I lived or died\n'Til I started stealing for my bread\nHey, hey, hey\nI've been down this road before\nAt least a hundred times or more\nEvery day it seems the same\nAnother dead end in the pouring rain\nSometimes you can't believe your eyes\nNobody warns you when you grow up\nJust what its like to be alive\nUnder the bridge, under the bridge\nI don't ever wanna die\nUnder the bridge, under the bridge\nI just wanna survive\nUnder the bridge, under the bridge\nUnder the bridge\nWhen I was younger I used to steal\nAnything I could get my hands on\nThen I found Jesus and I got real\nTurned my life around, I learned to sing the song\nHey, hey, hey\nDon't you think it's kind of funny\nHow we spend our lives in misery?\nEverybody wants to live so badly\nBut nobody seems to know the reason why\nSometimes you can't believe your eyes\nNobody tells you when you're growing up\nWhat it's really like to live your life\nUnder the bridge, under the bridge\nI don't ever wanna die\nUnder the bridge, under the bridge\nI just wanna survive\nUnder the bridge, under the bridge\nUnder the bridge",
		},
		{
			Band:   "Johnny Cash",
			Name:   "I Walk the Line",
			Lyrics: "You know I walk the line\nI keep a close watch on this heart of mine\n'Cause you're bound to reap just what you sow\nAnd I've sown some oats, now I'm reaping brome\nSo I'll just keep on walking that line\n\nI walk the line, I walk it\nFor my baby\nI'll walk the line for my sugar too\nBecause she's so good to me\nSaid she'd wait in old Galilee\nIf I ever wanted her to\nBut I keep on walking the line\n\nYou see, it's either love or hate\nThere's no in between\nWhen you play with fire, you're gonna get burned\nThat's a rule I always seen\nBut I'll keep on walking the line",
		},
		{
			Band:   "Johnny Cash",
			Name:   "Ring of Fire",
			Lyrics: "Love is a burning thing\nAnd it makes a fiery ring\nBound by wild desire\nI fell into a ring of fire\nAshes and embers grow\nMy one true passion so\nBurnin' within my soul\nThis ring of fire that I call my own\n\nWell, I courted a woman\nWho was tall and slim\nShe took my ring of fire and\nMade it a living dream\nWe lived in a mansion\nThough humble was our start\nNow we've gone through heaven\nAnd we've been through hell\nIn this ring of fire",
		},
		{
			Band:   "Bob Marley & The Wailers",
			Name:   "Get Up, Stand Up",
			Lyrics: "Stand up for your rights\nDon't give in to no one\nYou have the right to be free\nSo stand up for your freedom\n\nGet up, stand up\nStand up for your rights\nGet up, stand up\nDon't give up the fight\n\nThey say you can't do this\nSay you can't do that\nBut you know you've got the right\nTo decide your own fate\n\nGet up, stand up\nStand up for your rights\nGet up, stand up\nDon't give up the fight",
		},
		{
			Band:   "Bob Marley & The Wailers",
			Name:   "Redemption Song",
			Lyrics: "Old pirates, yes, they rob I\nSold I to the merchant ships\nMinutes after they took I\nFrom the bottomless pit\nBut my hand was made strong\nBy the hand of the Almighty\nWe forward in this generation\nTriumphantly\n\nWon't you help to sing\nThese songs of freedom?\n'Cause all I ever have\nRedemption song\nEmancipate yourselves from mental slavery\nNone but ourselves can free our minds\nHave no fear for atomic energy\n'Cause none of them can stop the time\nHow long shall they kill our prophets\nWhile we stand aside and look?\nSome say it's just a part of it\nWe've got to fulfill the book\n\nWon't you help to sing\nThese songs of freedom?\n'Cause all I ever had\nRedemption song",
		},
		{
			Band:   "Bob Marley & The Wailers",
			Name:   "One Love",
			Lyrics: "One love, one heart\nLet's get together and feel alright\nPeople, get ready\nFor the judgment day\nIt's the only way\nI'm gonna tell you now\nOne love, people get ready\n\n(One love)\n(People get ready)\n\nOne more river to cross\nOne more mountain to climb\nOne more burden to bear\nBefore I reach the other side\nGot to keep on pushing\nOnward to Zion\nGotta keep on moving\nMoving on\n\nOne love, one heart\nLet's get together and feel alright\nPeople, get ready\nFor the judgment day\nIt's the only way",
		},
	}
	return songs
}
