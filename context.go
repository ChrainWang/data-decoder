package decoder

type DecodeContext struct {
	Data        interface{}
	QueryKey    string
	QueryOption QueryOptionType

	done bool
}

func (ctx *DecodeContext) Done() { ctx.done = true }
