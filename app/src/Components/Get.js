import { useState } from "react";
import SetForm from "./Set";
import { get } from "../store";

function GetForm(props) {
    const [key, setKey] = useState('')
    const [isError, setIsErr] =  useState(false)
    const [errMsg, setErrMsg] = useState('')
    const [result, setResult] = useState('')
    const isInteger = (num) => {
        const number = Number(num);
        return Number.isInteger(number);
    }
    const handleForm = (e) => {
        const input = e.target.value
        if(!isInteger(input)) {
            setIsErr(true)
            setErrMsg("Input not an integer")
            return 
        }
        setIsErr(false)
        setErrMsg('')
        setKey(input)
        console.log(key)
    }
    const handleGet = async (e) => {
        setResult('')
        e.preventDefault()
        const data = await get(key, setIsErr, setErrMsg)
        //console.log("GOT Data : " + data.value)
        if (data && data.value) {
            setResult(data.value)
        }
        if (data && data.error) {
            setIsErr(true)
            setErrMsg(data.error)
        }
        setKey('')
        props.loadCache()
    }
    return (
        <div className="get-form border border-secondary border-2 rounded-3 p-4 m-2">
            <h6>Get Key</h6> 
            <div className="input-group mb-3">
                <input type="text" value={key} className="form-control" placeholder="Enter Key" aria-describedby="button-addon2" onChange={handleForm}/>
                <button onClick={handleGet}className="btn btn-outline-secondary" type="button" id="button-addon2">Get Value</button>
            </div>

            <div className="">
                <h6>Result</h6>
                {
                    isError ? <span className="badge bg-danger">{errMsg}</span> : ""
                }
                {
                    result !== '' ? <span className="badge bg-success"> Value : {result} </span> : ""
                }
              
            </div>

         </div>
    );
}

export default GetForm;