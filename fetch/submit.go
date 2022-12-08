package fetch

import (
	"fmt"
)

func Submit(
	day int,
	part1, part2 func(string) (string, error),
) (_, __ func(string) (string, error)) {
	return submissionWrapper(day, 1, part1),
		submissionWrapper(day, 2, part2)
}

func submissionWrapper(
	day, level int,
	in func(string) (string, error),
) func(string) (string, error) {
	return func(input string) (string, error) {
		answer, err := in(input)
		if err != nil {
			return ``, err
		}
		submitAnswer(day, level, answer)
		return answer, nil
	}
}

func submitAnswer(
	day, level int,
	answer string,
) {
	if hasCorrectAnswer(day, level) {
		return
	}
	if level == 2 && !hasCorrectAnswer(day, level) {
		return
	}

	var resp string
	fmt.Printf("Submit part %d answer? (Y/n) %q\n", level, answer)
	fmt.Scanf("%s", &resp)
	if len(resp) == 0 || (resp != `y` && resp != `Y`) {
		return
	}
	resp, err := postAnswerToWebsite(day, level, answer)
	if err != nil {
		fmt.Printf("error while submitting: %v\n", err)
		return
	}

	fmt.Printf("Successfully submitted: %q\n", resp)
	fmt.Printf("check the resp to find out if I got it right")
	/*
		<!DOCTYPE html>\n<html lang=\"en-us\">\n<head>\n<meta charset=\"utf-8\"/>\n<title>Day 8 - Advent of Code 2022</title>\n<!--[if lt IE 9]><script src=\"/static/html5.js\"></script><![endif]-->\n<link href='//fonts.googleapis.com/css?family=Source+Code+Pro:300&subset=latin,latin-ext' rel='stylesheet' type='text/css'/>\n<link rel=\"stylesheet\" type=\"text/css\" href=\"/static/style.css?30\"/>\n<link rel=\"stylesheet alternate\" type=\"text/css\" href=\"/static/highcontrast.css?0\" title=\"High Contrast\"/>\n<link rel=\"shortcut icon\" href=\"/favicon.png\"/>\n<script>window.addEventListener('click', function(e,s,r){if(e.target.nodeName==='CODE'&&e.detail===3){s=window.getSelection();s.removeAllRanges();r=document.createRange();r.selectNodeContents(e.target);s.addRange(r);}});</script>\n</head><!--\n\n\n\n\nOh, hello!  Funny seeing you here.\n\nI appreciate your enthusiasm, but you aren't going to find much down here.\nThere certainly aren't clues to any of the puzzles.  The best surprises don't\neven appear in the source until you unlock them for real.\n\nPlease be careful with automated requests; I'm not a massive company, and I can\nonly take so much traffic.  Please be considerate so that everyone gets to play.\n\nIf you're curious about how Advent of Code works, it's running on some custom\nPerl code. Other than a few integrations (auth, analytics, social media), I\nbuilt the whole thing myself, including the design, animations, prose, and all\nof the puzzles.\n\nThe puzzles are most of the work; preparing a new calendar and a new set of\npuzzles each year takes all of my free time for 4-5 months. A lot of effort\nwent into building this thing - I hope you're enjoying playing it as much as I\nenjoyed making it for you!\n\nIf you'd like to hang out, I'm @ericwastl on Twitter.\n\n- Eric Wastl\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n-->\n<body>\n<header><div><h1 class=\"title-global\"><a href=\"/\">Advent of Code</a></h1><nav><ul><li><a href=\"/2022/about\">[About]</a></li><li><a href=\"/2022/events\">[Events]</a></li><li><a href=\"https://teespring.com/stores/advent-of-code\" target=\"_blank\">[Shop]</a></li><li><a href=\"/2022/settings\">[Settings]</a></li><li><a href=\"/2022/auth/logout\">[Log Out]</a></li></ul></nav><div class=\"user\">(anonymous user #2702287) <span class=\"star-count\">15*</span></div></div><div><h1 class=\"title-event\">&nbsp;&nbsp;<span class=\"title-event-wrap\">{year=&gt;</span><a href=\"/2022\">2022</a><span class=\"title-event-wrap\">}</span></h1><nav><ul><li><a href=\"/2022\">[Calendar]</a></li><li><a href=\"/2022/support\">[AoC++]</a></li><li><a href=\"/2022/sponsors\">[Sponsors]</a></li><li><a href=\"/2022/leaderboard\">[Leaderboard]</a></li><li><a href=\"/2022/stats\">[Stats]</a></li></ul></nav></div></header>\n\n<div id=\"sidebar\">\n<div id=\"sponsor\"><div class=\"quiet\">Our <a href=\"/2022/sponsors\">sponsors</a> help make Advent of Code possible:</div><div class=\"sponsor\"><a href=\"https://www.tcgplayer.com/adventofcode/?utm_campaign=aoc&amp;utm_source=adventOfCode&amp;utm_medium=aocPromo\" target=\"_blank\" onclick=\"if(ga)ga('send','event','sponsor','sidebar',this.href);\" rel=\"noopener\">TCGplayer</a> - Join an ever-growing worldwide team in connecting hobbyists to communities, exchanging things and thoughts that fuel passions. Competitive benefits, remote work, participate in hackathons and Advent of Code, &amp; more!</div></div>\n</div><!--/sidebar-->\n\n<main>\n<article><p>That's the right answer!  You are <span class=\"day-success\">one gold star</span> closer to collecting enough star fruit. <a href=\"/2022/day/8#part2\">[Continue to Part Two]</a></p></article>\n</main>\n\n<!-- ga -->\n<script>\n(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){\n(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),\nm=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)\n})(window,document,'script','//www.google-analytics.com/analytics.js','ga');\nga('create', 'UA-69522494-1', 'auto');\nga('set', 'anonymizeIp', true);\nga('send', 'pageview');\n</script>\n<!-- /ga -->\n</body>\n</html>
		<!DOCTYPE html>\n<html lang=\"en-us\">\n<head>\n<meta charset=\"utf-8\"/>\n<title>Day 8 - Advent of Code 2022</title>\n<!--[if lt IE 9]><script src=\"/static/html5.js\"></script><![endif]-->\n<link href='//fonts.googleapis.com/css?family=Source+Code+Pro:300&subset=latin,latin-ext' rel='stylesheet' type='text/css'/>\n<link rel=\"stylesheet\" type=\"text/css\" href=\"/static/style.css?30\"/>\n<link rel=\"stylesheet alternate\" type=\"text/css\" href=\"/static/highcontrast.css?0\" title=\"High Contrast\"/>\n<link rel=\"shortcut icon\" href=\"/favicon.png\"/>\n<script>window.addEventListener('click', function(e,s,r){if(e.target.nodeName==='CODE'&&e.detail===3){s=window.getSelection();s.removeAllRanges();r=document.createRange();r.selectNodeContents(e.target);s.addRange(r);}});</script>\n</head><!--\n\n\n\n\nOh, hello!  Funny seeing you here.\n\nI appreciate your enthusiasm, but you aren't going to find much down here.\nThere certainly aren't clues to any of the puzzles.  The best surprises don't\neven appear in the source until you unlock them for real.\n\nPlease be careful with automated requests; I'm not a massive company, and I can\nonly take so much traffic.  Please be considerate so that everyone gets to play.\n\nIf you're curious about how Advent of Code works, it's running on some custom\nPerl code. Other than a few integrations (auth, analytics, social media), I\nbuilt the whole thing myself, including the design, animations, prose, and all\nof the puzzles.\n\nThe puzzles are most of the work; preparing a new calendar and a new set of\npuzzles each year takes all of my free time for 4-5 months. A lot of effort\nwent into building this thing - I hope you're enjoying playing it as much as I\nenjoyed making it for you!\n\nIf you'd like to hang out, I'm @ericwastl on Twitter.\n\n- Eric Wastl\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n-->\n<body>\n<header><div><h1 class=\"title-global\"><a href=\"/\">Advent of Code</a></h1><nav><ul><li><a href=\"/2022/about\">[About]</a></li><li><a href=\"/2022/events\">[Events]</a></li><li><a href=\"https://teespring.com/stores/advent-of-code\" target=\"_blank\">[Shop]</a></li><li><a href=\"/2022/settings\">[Settings]</a></li><li><a href=\"/2022/auth/logout\">[Log Out]</a></li></ul></nav><div class=\"user\">(anonymous user #2702287) <span class=\"star-count\">16*</span></div></div><div><h1 class=\"title-event\">&nbsp;<span class=\"title-event-wrap\">{&apos;year&apos;:</span><a href=\"/2022\">2022</a><span class=\"title-event-wrap\">}</span></h1><nav><ul><li><a href=\"/2022\">[Calendar]</a></li><li><a href=\"/2022/support\">[AoC++]</a></li><li><a href=\"/2022/sponsors\">[Sponsors]</a></li><li><a href=\"/2022/leaderboard\">[Leaderboard]</a></li><li><a href=\"/2022/stats\">[Stats]</a></li></ul></nav></div></header>\n\n<div id=\"sidebar\">\n<div id=\"sponsor\"><div class=\"quiet\">Our <a href=\"/2022/sponsors\">sponsors</a> help make Advent of Code possible:</div><div class=\"sponsor\"><a href=\"https://www.retool.com\" target=\"_blank\" onclick=\"if(ga)ga('send','event','sponsor','sidebar',this.href);\" rel=\"noopener\">Retool</a> - Quickly build CRUD apps, admin panels, recurring jobs, and dashboards with JavaScript. Run in our cloud or yours. Free for small teams.</div></div>\n</div><!--/sidebar-->\n\n<main>\n<article><p>That's the right answer!  You are <span class=\"day-success\">one gold star</span> closer to collecting enough star fruit.</p><p>You have completed Day 8! You can <span class=\"share\">[Share<span class=\"share-content\">on\n  <a href=\"https://twitter.com/intent/tweet?text=I+just+completed+%22Treetop+Tree+House%22+%2D+Day+8+%2D+Advent+of+Code+2022&amp;url=https%3A%2F%2Fadventofcode%2Ecom%2F2022%2Fday%2F8&amp;related=ericwastl&amp;hashtags=AdventOfCode\" target=\"_blank\">Twitter</a>\n  <a href=\"javascript:void(0);\" onclick=\"var mastodon_instance=prompt('Mastodon Instance / Server Name?'); if(typeof mastodon_instance==='string' && mastodon_instance.length){this.href='https://'+mastodon_instance+'/share?text=I+just+completed+%22Treetop+Tree+House%22+%2D+Day+8+%2D+Advent+of+Code+2022+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2022%2Fday%2F8'}else{return false;}\" target=\"_blank\">Mastodon</a\n></span>]</span> this victory or <a href=\"/2022\">[Return to Your Advent Calendar]</a>.</p></article>\n</main>\n\n<!-- ga -->\n<script>\n(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){\n(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),\nm=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)\n})(window,document,'script','//www.google-analytics.com/analytics.js','ga');\nga('create', 'UA-69522494-1', 'auto');\nga('set', 'anonymizeIp', true);\nga('send', 'pageview');\n</script>\n<!-- /ga -->\n</body>\n</html>
	*/
}

func hasCorrectAnswer(
	day, level int,
) bool {
	return false
}

func recordCorrectAnswer(
	day, level int,
	answer string,
) {
	// TODO
}
