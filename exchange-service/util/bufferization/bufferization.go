package bufferization

const(
	bufferSize = 50
)

func DoWithBufferization(task func([]string) (interface{}, error), params []string) (interface{},error) {

	var result interface{}
	var err error

	for len(params) > 0 {
		bufferedParams := make([]string, 0)

		if len(params) > bufferSize {
			bufferedParams = params[:bufferSize]
			params = params[bufferSize:]
		} else {
			bufferedParams = params
			params = params[len(params):]
		}

		result, err = task(bufferedParams)
		if err != nil{
			return nil, err
		}
	}

	return result,nil
}