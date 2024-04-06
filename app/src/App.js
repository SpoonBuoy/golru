import { useEffect, useState } from 'react';
import CacheView from './Components/CacheView';
import GetForm from './Components/Get';
import SetForm from './Components/Set';
import { getEntries } from './store';

function App() {
    const [entries, setEntries] = useState([])
    const loadCache = async () => {
        //it will load cache from backend
        console.log("Load Cache Called")
        const data = await getEntries()
        //console.log(data.entries)
        if (data && data.entries) {
          setEntries(data.entries)
        }
    }
    useEffect(() => {
      const data = getEntries()
      if (data && data.entries) {
        setEntries(data.entries)
      }
    }, [])

  return (
    <div className="App">
      <h1 className="text-center"> LRU Cache </h1>
      <GetForm loadCache={loadCache}/>
      <SetForm loadCache={loadCache}/>
      <CacheView entries={entries}/>
    </div>
  );
}

export default App;
