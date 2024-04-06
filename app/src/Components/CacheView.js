import { useState } from "react";
import SetForm from "./Set";

function CacheView(props) {
    
    return (
        <div className="get-form border border-secondary border-2 rounded-3 p-4 m-2">
            <h6> Top 10 Entries in LRU </h6>
            <table class="table">
            {
                props.entries && props.entries.length > 0 
                ? <thead>
                    <tr>
                    <th scope="col">Key</th>
                    <th scope="col">Value</th>
                    <th scope="col">Expiry</th>
                    </tr>
                </thead>
                :""
            }
            <tbody>
            {
                props.entries.map((e) => {
                    return <tr key={e.key}>
                        <td> {e.key} </td>
                        <td> {e.value} </td>
                        <td> {e.expiry} </td>
                    </tr>
                })
            }
            </tbody>
        </table>
         </div>
    );
}

export default CacheView;