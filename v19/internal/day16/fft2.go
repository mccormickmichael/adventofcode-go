package day16


/*
 Cheater method of calculating the phase results of a message with an offset indicator:
 Values calculated past the offset point do not rely on ANY values calculated prior to the offset point.
 If the offset is at least halfway through the message, then the base pattern prior to the offset
 is all zeroes, and the base pattern after the offset is all ones. So the phase value of elements beyond
 the offset is the sum of values beyond that element.

 We can optimize this calculation by working backwards: the value of a row is N + sum(N+1:), but sum(N+1:)
 is just the sum of the next row. So starting at the end, work backwards to the offset, assigning values as
 we accumulate them.
 */
func fft2(msg []int) {
	accumulator := 0
	for i := len(msg)-1; i >= 0; i-- {
		accumulator += msg[i]
		msg[i] = accumulator % 10
	}
}