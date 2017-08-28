package main

func Boolfuck(code, input string) string {
	
	is := []byte{}      // input stack
	os := []byte{}      // output stack
	ms := []byte{0}     // memory stack
	
	cp := 0             // code position
	ip := 0             // input position
	mp := 0             // memory position

	bp := uint8(0)      // bit position
	
	// Convert the input string into a binary input stream (is)
	for _, i := range input {
		if i > 255 {
			i = 255
		}
		
		b := byte(i)
		m := byte(1)

		for j := 0; j < 8; j++ {
			if b & m > 0 {
				is = append(is, 1)
			} else {
				is = append(is, 0)
			}
			m *= 2
		}
	}

	// Iterate over the program code, advancing with the code counter (cp)
	for {
		switch code[cp] {
			case '+':
				if ms[mp] == 0 {
					ms[mp] = 1
				} else {
					ms[mp] = 0
				}

			case ',':
				if ip >= len(is) {
					ms[mp] = 0
				} else {
					ms[mp] = is[ip]
					ip++
				}

			case ';':
				// bit position (bp) determines which bit in the output stream (os) should be set
				if bp == 0 {
					os = append(os, 0)
				}
				
				os[len(os)-1] |= (ms[mp] << bp)
				
				bp++
				if(bp > 7) {
					bp = 0
				}

			case '<':
				mp--
				// Prepend a new byte to the beginning of the memory stack.
				if(mp < 0) {
					mp = 0
					ms = append([]byte{0}, ms...)
				}

			case '>':
				mp++
				// Append a new byte to the end of the memory stack.
				if mp >= len(ms) {
					ms = append(ms, 0)
				}

			case '[':
				if ms[mp] == 0 {
					// Keep track of the bracket depth (bd)
					bd := 1

					// Increment the code counter (cp) until a matching closing bracket is found
					for {
						cp++

						if code[cp] == '[' {
							bd++
						} else if code[cp] == ']' {
							bd--
							if bd < 1 {
								break
							}
						}
					}
				}

			case ']':
				// Keep track of the bracket depth (bd)
				bd := 1

				// Decrement the code counter (cp) until a matching opening bracket is found.
				for {
					cp--

					if code[cp] == ']' {
						bd++
					} else if code[cp] == '[' {
						bd--
						if bd < 1 {
							cp--
							break
						}
					}
				}
		}

		// Increment the code counter (cp). Stop execution if the counter exceeds the program length.
		cp++
		if cp >= len(code) {
			break
		}
	}

	return string(os)
}