import React, { useState } from 'react'
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '../../components/ui/card'
import { Label } from '../../components/ui/label'
import { Input } from '../../components/ui/input'
import { Button } from '../../components/ui/button'
import { createUserWithEmailAndPassword, getAuth } from 'firebase/auth'
import { firebaseApp } from '../../auth/context/auth'

function RegisterView() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSignUp = () => {
    const auth = getAuth(firebaseApp);
  createUserWithEmailAndPassword(auth, email, password)
    .then((userCredential) => {
      const user = userCredential.user;
      console.log(user);
    })
    .catch((error) => {
      console.log(error);
         
    });
  }

  return (
    <div className='flex justify-center items-center h-screen w-full'>
      <Card className="w-[80%]">
      <CardHeader className='text-start'>
        <CardTitle>Register</CardTitle>
        <CardDescription>Enter your detailes</CardDescription>
      </CardHeader>
      <CardContent className='text-start'>
        <form>
          <div className="grid w-full items-center gap-4">
          <div className="flex flex-col space-y-2">
              <Label htmlFor="name">Name</Label>
              <Input id="name" placeholder="Enter your Name" type='text' />
            </div>
            <div className="flex flex-col space-y-2">
              <Label htmlFor="name">Email</Label>
              <Input id="email" placeholder="Enter your Email" type='email' onChange={(e) => setEmail(e.target.value)} />
            </div>
            <div className="flex flex-col space-y-2">
              <Label htmlFor="name">Password</Label>
              <Input id="password" placeholder="Enter your Password" type='password' onChange={(e) => setPassword(e.target.value)} />
            </div>
          </div>
        </form>
        
      </CardContent>
      <CardFooter className="flex justify-end ">
        <Button onClick={handleSignUp}>SignUp</Button>
      </CardFooter>
    </Card>
    </div>
  )
}

export default RegisterView
