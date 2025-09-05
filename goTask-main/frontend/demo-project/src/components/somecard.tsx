import "../output.css"
import type { User } from "../models/user";
import P from "./pcomp"
import { useState } from "react";



type UserCardProps = {
    user: User;
};

function SomeCard({ user }: UserCardProps) {
    const [count, setCount]=useState<number>(0);

    function handleClick(){
            setCount(onSetCount);
    }

    function onSetCount(c:number){
       return c+1;
    }

    return (
        <>
            <div className="mx-auto flex max-w-sm items-center gap-x-4 rounded-xl bg-white p-6 shadow-lg outline outline-black/5 dark:bg-slate-800 dark:shadow-none dark:-outline-offset-1 dark:outline-white/10">
                <img className="size-12 shrink-0" src="/vite.svg" alt="ChitChat Logo" />
                <div>
                    <div className="text-xl font-medium text-black dark:text-white">Details</div>
                    <P data={user.name}></P>
                    <P data={user.email}></P>
                    <P data={user.contact}></P>
                    {/* <button onClick={handleClick}>count is {count}</button> */}
                    <button onClick={()=>setCount((count)=>count+1)}>count is {count}</button>
                </div>
            </div>
        </>
    )
}

export default SomeCard;

// npx @tailwindcss/cli -i ./src/tailwind.css -o ./src/output.css --watch
