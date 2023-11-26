import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import axios from 'axios'
import './App.css'
// import Message from './Message'

type YoutubeData = {
  contentDetails: {
    relatedPlaylists: {
      uploads: string;
    };
  };
  etag: string;
  id: string;
  kind: string;
  snippet: {
    customUrl: string;
    description: string;
    localized: {
      description: string;
      title: string;
    };
    publishedAt: string;
    thumbnails: {
      default: {
        height: number;
        url: string;
        width: number;
      };
      high: {
        height: number;
        url: string;
        width: number;
      };
      medium: {
        height: number;
        url: string;
        width: number;
      };
    };
    title: string;
  };
  statistics: {
    subscriberCount: string;
    videoCount: string;
    viewCount: string;
  };
}

const App = () => {
  const [count, setCount] = useState(0)
  const [msg, setMsg] = useState<YoutubeData[]>([])

  useEffect(() => {
    (async () => {
      console.log("aaa")
      try {
        const res = await axios.get('http://localhost:8080/api')
        console.log(res.data)
        setMsg(res.data)
        // console.log(res.data)
      } catch(e) {
        console.log(e)
      }
    })()
  },[])

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
          <ul>
            {msg.map((data) => (
              <li key={data.id}>{data.id}</li>
            ))}
          </ul>
        <p>
          {/* <Message messages={msg}/> */}
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
