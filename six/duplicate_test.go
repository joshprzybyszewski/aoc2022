package six

import "testing"

var (
	input = "jfnjjwbbqttplpvllqgllmdllfmllscssqmqzmmwzznqnwqnwnqnjjbdbpbtbdbzzzljljzjjpccrmmppzfpzfpfnfccfbbcqcrcffblfbftbfbtbwwwmgwmgmnngnllnfllhghcghhjppchcfcnfffllmmqbmmpwwwwlqwwqgqcqsqjqpqzqqdzdtztltslsljjfqfcqqgbqqqghqgqvgvrggqwggrgjgmmnrmmzgmzgzpzjjctcmtcmcnndppcvpvrrwvrvhrvhrhjhnjnvjnjrjggccvffnqqvfqvvnmvmqmfmfqqzfzbfzzzgpzpllrwwnpwpnwnwgwhhrrdnrrdjjzjszsjjbddcdbbvmbmqbqnbqbsqbsqqwbwhwggssdnnmttvnnvmnmhmfhhjchcttzdzdqqszzcwwhhwzhwhphqhcqqsggddfmmvzmzwmwfwzwrrbmrrnwnfnwnlwnwrwfwnnmtnnzwnwdnnbhhrphrhlhwllpmmbcbtbffmqffddjnjwwzpzfpptbbqqwbwzbzjbjmjljblbtlblqqhqbqggrngrgllbmbccmhmqmqwwqcqssqzzfjzjrjnnqrqssfnsnvvtgvvmsvsqqljjbsbrrjllvfvzfzmzhzzhthjhshlslfljfjqjpqpvvmpmhpmhmqqmmdwmddppjlplhlsstlssgnggrbblggffcdfdzzwqqtztqtwqtwtzzsjsbszsbsvbvwwjqjnnpdpccwssvdsdzzqbqbtbtqtmtltltvlvddzwzzfpzpjpgphprpgpqqwppdwpdplddvffcdffvpvqqgvqgvvrfvrvqrrcjcpjjpttftqqvjqvqsvqsvqssdpdbbbmcmscsddbhhgttwhhjlltqllnqntqtsscnntwwhswswlwggldltlttsjszsnznsznzccbtbblplnnmfmqmrrvjvhjhzhnzzgnnhrrdrllblpbllfdfjjssvnssvlsllnqqhwqhhhsgstsjstthrrhrghrhhfmhmwhwrwwsrwrfwwdnntqnnsvnvmnnfvnntztqzqhqnqjnnjfflbfllrsllqhqdqccgvgnvvcwcfccmssqnqhhqrrfrtrvvnjnpjnjjpplmlppvmpphjhppvhvdvssjcjrrtrdrrsvvbbjzzrtztgzghzhccwmccshhzbhhdwdwsdswwlcwllpblpphrppfhfnffrbbcgcmggnvnzzmvvcrrftrftrffcscvcsslbljlglzgzbzczszmsmbmnbbhdhvvsqvqhvvfrfddbpwgvztwwqcpzhhwnhphnrwldjmztsptbbgsqbqqccwbdqzvhfjlfldgphzbfprclgpfztbrgvsvfpghmdchscbdqjqgzvmrtdrfzbhgdvgznjcsmglcfwhdtpsljnvvzjcbbrczwtgpdmgpzhctvbbmvsjzthffsjqhfsdrclpqslbhnmpczwvggpzbjcchfjzjhhgtrmlgnzlndfvzrccgggrpmprbmjbfjjhzrhrtwgqdbgdlqghssrnmtmpvttcqwnwdzhgfnddgbqcsdvzvwqdnmmpwrwhfbqtcpqhvwbczrmjqzsntvdrncwjsmvvwcngrtlwtjmnctwrrtvphbjhlqmgzfsfsrblzzvmzlbhzjhwbdfpncdrfchmrqhspdszcjrnvwtmjzmsmzcdphsdzjgqswwrpdvlpvrdnhplnlmswvcrzlcmbtqtscjfwrnrctrvdqcqzwcvgvpdgrndrgsrvzftwpqjjgjhzwhvrjlqntdtcjdrqzhqlqqdffcgvttlhvwgggnwmdlvghfgjpsmntbvbjbbttrwsljwsrvtmznvqdptpwtdcwtcsfdjlmdqthqggjcptrqhbsbjzqqmvvjmgmppqmjmnjdqvspzlbgzjsjshpslmszqnzghsszpsmpzfcrqqjdwvtbnzstvvjzvtzgpptcmvmbvmpvpzvgfnwtlmdzhvhshtwvnbgwmtzqhcptflpqsqvmptchpfcbwhvjzdcnsnqrgdwfcthqfssnbqnvgvvhlzqfqmdlcwnshtvhhhpghjbmhdbfbqcvbnbvwbzcbbmjnrqmsdqnmnbsrvhggzsrlbwtfmgwrnlhrbrrrqdcspnrpnppngrtdqtbmbhcbjrlhpfjpdnfndmqvwvhlgmsntpwrlrwwqhwvzbpzqqggnbqlsjjqtbqjcdpmndgmtdhfbqrpdzzsnmhzmqqnbdqftqmnhfbdzdlfwgjsjhrcsmtfzgwbvbbzdrlbmcgmppqfppmbqrnsmrmhrdsvgcfmzpfnvrbbgfccfcbphszwdbnnwcjjvvlpdtfzgtslvgqwmsvlpzjcbqwqclrjrsgthhtqrqrhvsdfjntgllsvslrvdtnsdmrgtqcmswnqwlrwlfmcfftbjpvdnmczqzldsssszhjtqtqvqtwhjcqchjvqvntvzzzprbmjcctsqfdcvpbtsgnnsqtqnmjhrgqcjnzrdsgrbtdpqjbgcmnfwhnsrfwcdmncjzwcngfbmmrsbvgvvqpvrdjfsqwjdmqjdpzcbjjfmzjjgbnwqgrvpmbzdhsgtldrzvglscfwbmjltcrzrgdslgprwscwbrhtdtglznjdcvfjzjjqzntdqdbcrcbbmvnzdshjzcsfsgpghmgdqdwsnwjtvtbqbqccbcwjpnhdhzcvdssvnvqtvzwprhpgftdwwvgsbnlzzjppcrrwmrsthvjjrvrsdrbdqfgsjsmwfplpstrbnpdhhcblhjfwzngmhlwbvnfcbgwshspsbbgbldrvmcnczszpgnddrfwrtgcqjggrrcbjwrdjlrvtspbftrtjbzjwchpfnjctcjtwtpmtblczcftqlphdjczfrvtzlsglpvhqsqqblttdjrlczhrqsgpggmvnhpqtrfbpgvzftwtsmwhwswtpvtwnsshmlcffpcjshqhqqsjtpbgszscmcbnhjjtjmpgfdhgmljqmmwlfptstjjvqhcbjpjpwzwqflhslclzzjlmcttbsncqmfzhgnzwbdtnvfwbtztwbhtfsqjfzwmfflmbwnqzqhcjwdpbvngsgzlwvwcqhqjsndznbbdcqqhmjjpqjbsnvwztgmqwdcbbjvcndmhsbvbjnzlbscmgnjcrrwrfdljtcsgmwtffgcjflpzzdcnzvmrbnrjbbmhzqqjtgsrwqmmrhpndwlbnrtrhhpqlmdrcrtdmzsslrmffpftdjvfcpvvhzhjhqtrrsclvtbsccgmmqrjbqgbmpnbzlsncssdhmjppjptvddfgbbnjzjjldjlqjzhhttsclrmsgzctwjqqvtjlfzwgtffgrdjzwdcnrprlcswffghngrqcgsbzqhhvbfjtwcjlrrmbtqjdrgpnbftnmzqnndnqwgrqndlwmjnnspbhjlnzrnptnrmcjhpbfcqpvbchvdwthjlcrfpssgtfbsgfrftcrwttrspbsvzpvcczmdqslcdgfljvtjsdpjnwmdvfzfllrdrbgvpltzlqcrlwbncswhfvrdthspmhfhfdlvpbcqlmjfznhnqblffftgzqrtswnmtnvjprqqhhhvrscvbbzgmnlnprghfdjqbgjppjzjrnclfdssbmgspwcscnlcrrqmtlljrmcwgdgcqwvvjzvsjdjvsspszlcthwzrwqtzdgmqvnlvvzrvrpqqwswzcchncrpnjdmflvmhhwvrrstpvnszfrmvpdtpqpbdmwvvbbpjnwmtststtlcvqdnvqqphzlhhzbbbjssgdcnhlmwrzwvwmcmgrcngqzcnffqzfnvldpdjmsspgpbrzhnszfnljfcrgsjvqjjbstvghlcslhqlzhltpglwffrzfgjghssfgrptbnpbhqnhhfbjsnmsvltqpthdmzzrhrhhmzlplvrtdqfrfrppdpqnllblcfjqpdwznsbrhcncdpmztcrjrfnlwtznrmpbzqsbrqrbnthgfpshrdhnwjmrnsmsfqwdjsmsvhfrbdpjrwcvmdvvmdtfqjgmdsrqtctsdmznngbsrfjvhllgwt\n"
)

func BenchmarkFind(b *testing.B) {
	/*
		goos: linux
		goarch: amd64
		pkg: github.com/joshprzybyszewski/aoc2022/six
		cpu: Intel(R) Core(TM) i5-3570 CPU @ 3.40GHz
		BenchmarkFind/One-4         	  360752	      3122 ns/op	       4 B/op	       1 allocs/op
		BenchmarkFind/Two-4         	  435723	      2796 ns/op	       4 B/op	       1 allocs/op
		PASS
		ok  	github.com/joshprzybyszewski/aoc2022/six	3.410s
	*/

	b.Run(`One`, func(b *testing.B) {
		var answer string
		var err error
		for n := 0; n < b.N; n++ {
			answer, err = One(input)
			if answer != `1300` {
				b.Fail()
			}
			if err != nil {
				b.Fail()
			}
		}
	})

	b.Run(`Two`, func(b *testing.B) {
		var answer string
		var err error
		for n := 0; n < b.N; n++ {
			answer, err = Two(input)
			if answer != `3986` {
				b.Fail()
			}
			if err != nil {
				b.Fail()
			}
		}
	})
}
