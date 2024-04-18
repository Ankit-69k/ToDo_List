import './App.css'
import HeadCard from './components/headCard'
import { DataTable } from './components/todoTable'
import { columns } from "./components/column"


function App() {
  return (
    <div className="h-full flex-1 flex-col space-y-8 p-8 md:flex">
      <HeadCard /> 
    </div>
  )
}

export default App
