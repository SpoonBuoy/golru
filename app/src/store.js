const port = '9000'
const api = 'http://localhost:' + port 

export const set = async (key, value, expiry, setIsErr, setErrMsg) => {
try {
    const response = await fetch(api, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        key: parseInt(key),
        value: parseInt(value),
        expiry: parseInt(expiry)
      }),
    })

    if (!response.ok) {
        setIsErr(true)
        setErrMsg("Error")
      throw new Error('Network response was not ok')
    }

    const data = await response.json()
    console.log(data)
  } catch (error) {
    setIsErr(true)
    setErrMsg("Error : " + error)
    console.error('Error : ', error)
  }
}
export const get = async(key, setIsErr, setErrMsg) => {
    try {
        const response = await fetch(api + '/' + key)

        if (!response.ok) {
            setIsErr(true)
            setErrMsg("Error")
            throw new Error('Network response was not ok')
        }

        const data = await response.json()
        return data
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error)
    }
}
export const getEntries = async(key, setIsErr, setErrMsg) => {
    try {
        const response = await fetch(api + '/top')

        if (!response.ok) {
            setIsErr(true)
            setErrMsg("Error")
            throw new Error('Network response was not ok')
        }

        const data = await response.json()
        return data
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error)
    }
}