import './App.css'
import HeadCard from './components/headCard'
import { PaginationDemo } from './components/pagination'
import { DataTableDemo } from './components/todoTable'

function App() {
  return (
    <div className=" h-full flex-1 flex-col space-y-8 p-8 md:flex">
      <HeadCard />
      <DataTableDemo />
      <PaginationDemo />
    </div>
  )
}

export default App
