package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var res Resource

	for range 100 {
		res, err = opener()
		_, isTransient := err.(TransientError)

		if err != nil && !isTransient {
			return err
		}

		if err == nil {
			break
		}
	}

	defer func(res Resource, input string) {
		if rec := recover(); rec != nil {
			err = rec.(error)

			if errFrob, is := rec.(FrobError); is {
				res.Defrob(errFrob.defrobTag)
				err = errFrob
			}
		}

		res.Close()
	}(res, input)

	res.Frob(input)

	return
}
