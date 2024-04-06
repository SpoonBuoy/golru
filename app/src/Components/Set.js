import { useState } from "react";
import { get, set } from "../store";

function SetForm(props) {
    const [key, setKey] = useState('')
    const [value, setValue] = useState('')
    const [expiry, setExpiry] = useState('')
    const [isError, setIsErr] =  useState(false)
    const [isSubmitted, setIsSubmitted] = useState(false)
    const [errMsg, setErrMsg] = useState('')
    const [buttonText, setButtonText] = useState('Submit')
    const isInteger = (num) => {
        const number = Number(num);
        return Number.isInteger(number);
    }
    const handleKey = (e) => {
        const input = e.target.value
        setKey(input)
        console.log(input)
        if(!isInteger(input)) {
            setIsErr(true)
            setErrMsg("Key must be an integer")
            return 
        }
        setIsErr(false)
        setErrMsg('')
       // console.log(key)
    }
     const handleValue = (e) => {
        const input = e.target.value
        setValue(input)
        console.log(input)
        if(!isInteger(input)) {
            setIsErr(true)
            setErrMsg("Value must be an integer")
            return 
        }
        setIsErr(false)
        setErrMsg('')
       // console.log(key)
    }
     const handleExpiry = (e) => {
        const input = e.target.value
        setExpiry(input)
        console.log(input)
      
        if(!isInteger(input)) {
            setIsErr(true)
            setErrMsg("Expiry must be an integer")
            return 
        }
        setIsErr(false)
        setErrMsg('')
       // console.log(key)
    }
    const handleSubmit = async (e) => {
        e.preventDefault()
        if (key.length === 0 || value.length === 0 || expiry.length === 0){
            setErrMsg("All fields are required")
            return
        }
        if(isError) {
            setErrMsg("Please Fix all the errors before submitting")
            return
        }
        await set(key, value, expiry, setIsErr, setErrMsg)
        console.log(isError, isSubmitted)
        setIsSubmitted(true)
        setKey('')
        setValue('')
        setExpiry('')
        props.loadCache()
    }
    return (
        <div className="get-form border border-secondary border-2 rounded-3 p-4 m-2">
            <h6>Set</h6> 
                {
                    isError ? <span className="badge bg-danger">{errMsg}</span> : <span></span>
                }
              
            <form>
                <label for="key" className="form-label">Key </label>
                <input onChange={handleKey} value = {key} id="key" type="text" className="form-control" placeholder="Enter Value (Integer Only)"></input> <br></br>
                <label for="value" className="form-label">Value </label>
                <input onChange={handleValue} value = {value} id="value" type="text" className="form-control" placeholder="Enter Value (Integer Only)"></input> <br></br>
                <label for="expiry" className="form-label">Expiry </label>
                <input onChange={handleExpiry} value = {expiry} id="expiry" type="text" className="form-control" placeholder="Expiry in Seconds (Integer Only)"></input> <br></br>
            <button onClick={handleSubmit} className="btn btn-primary">{buttonText}</button>
            </form> 

         </div>
    );
}

export default SetForm;