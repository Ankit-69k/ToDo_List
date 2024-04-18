import './App.css'
import HeadCard from './components/headCard'
import { DataTable } from './components/todoTable'
import { columns } from "./components/column"
import { Tasks } from './data/task';


function App() {
  return (
    <div className="h-full flex-1 flex-col space-y-8 p-8 md:flex">
      <HeadCard /> 
      <DataTable data={Tasks} columns={columns} />
    </div>
  )
}

export default App
