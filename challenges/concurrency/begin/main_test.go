package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func BenchmarkDoAnalysis(b *testing.B) {
	reader := strings.NewReader(data)
	for i := 0; i < b.N; i++ {
		doAnalisys(reader, analysts()...)
		reader.Seek(0, io.SeekStart)
	}
}

func TestDoAnalysis(t *testing.T) {
	reader := strings.NewReader(data)
	analysts := analysts()
	doAnalisys(reader, analysts...)
	for _, analist := range analysts {
		fmt.Printf("%s: %d", analist.name(), analist)
	}
}

const data = `
When has justice ever been as simple as a rule book?
Worf, It's better than music. It's jazz. Your head is not an artifact! Mr. Worf, you sound like a man who's asking his friend if he can start dating his sister. My oath is between Captain Kargan and myself. Your only concern is with how you obey my orders. Or do you prefer the rank of prisoner to that of lieutenant? Ensign Babyface! and attack the Romulans. I guess it's better to be lucky than good. That might've been one of the shortest assignments in the history of Starfleet. But the probability of making a six is no greater than that of rolling a seven. You did exactly what you had to do. You considered all your options, you tried every alternative and then you made the hard choice. Your shields were failing, sir. Fate. It protects fools, little children, and ships named "Enterprise." You enjoyed that. I've had twelve years to think about it. And if I had it to do over again, I would have grabbed the phaser and pointed it at you instead of them. Shields up! Rrrrred alert! Some days you get the bear, and some days the bear gets you. Captain, why are we out here chasing comets? This is not about revenge. This is about justice. Not if I weaken first. Computer, lights up! Is it my imagination, or have tempers become a little frayed on the ship lately? Fear is the true enemy, the only enemy. The look in your eyes, I recognize it. You used to have it for me. For an android with no feelings, he sure managed to evoke them in others. Mr. Worf, you do remember how to fire phasers? I'd like to think that I haven't changed those things, sir. Then maybe you should consider this: if anything happens to them, Starfleet is going to want a full investigation. Now, how the hell do we defeat an enemy that knows us better than we know ourselves? Sorry, Data. We finished our first sensor sweep of the neutral zone. They were just sucked into space. I suggest you drop it, Mr. Data. The unexpected is our normal routine. And blowing into maximum warp speed, you appeared for an instant to be in two places at once. Fate protects fools, little children and ships named Enterprise. We could cause a diplomatic crisis. Take the ship into the Neutral Zone When has justice ever been as simple as a rule book? Is the meaning of life 42?
`