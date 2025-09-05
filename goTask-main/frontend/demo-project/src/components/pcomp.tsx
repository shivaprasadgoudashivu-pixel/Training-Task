
type Prop = {
  data?: string;
};

export default function P({data}:Prop){
    return(
        <>
        <p className="text-gray-500 dark:text-gray-400">{data}</p>
        </>
    )
}