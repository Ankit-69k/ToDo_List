import HeadCard from '../../components/headCard'
import { DataTable } from '../.././components/todoTable'
import { columns } from "../../components/column"
import { useEffect } from 'react';
import axios from 'axios';
import React from 'react';
import { ToDoItem } from '../../types/type';
import { Schema } from '../../data/schema';
import { priorities } from '../../data/data';
import { useAuthContext } from '../../hooks/auth';


function HomeView() {
    const [Tasks, setTasks] = React.useState([]);
    const { user } = useAuthContext();
  useEffect(() => {
    axios.get(`http://localhost:9000/api/tasks/${user?.uid}`)
      .then(response => {
        console.log(response.data);
        const filteredTasks = response.data.map((task : Schema,i:number) => ({
          id: task.id,
          index:i+1,
          title: response.data[i].task.title,
          status: response.data[i].task.status, 
          priorities: "low",
        }));
        setTasks(filteredTasks);
      })
      .catch(error => console.error(error));
  }, []);
  
  return (
    <div className="h-full flex-1 flex-col space-y-8 p-8 md:flex">
      <HeadCard /> 
      <DataTable data={Tasks} columns={columns} />
    </div>
  )
}

export default HomeView
