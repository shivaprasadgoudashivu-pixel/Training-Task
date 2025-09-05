import Demo from './components/demo.tsx'
import Logo from './components/logo.tsx'
import Card from './components/card.tsx'
import Misc from './components/misc.tsx'
import './App.css'
import SomeCard from './components/somecard.tsx'
import Header  from './components/header.tsx'
import type {User} from './models/user.ts'

function App() {
    const user: User = {
    contact: "9618558500",
    name: "Jiten Palaparthi",
    email: "jiten@example.com",
  };

  return (
    <>
    <Header></Header>
     <Logo></Logo>
     <Demo></Demo>
     <Card></Card>
     <Misc></Misc>
     <SomeCard user={user}/>
    </>
  )
}

export default App
