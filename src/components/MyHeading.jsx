// const MyHeading =()=> <h1>Heading</h1>  //to make code short if you dont use return then you can type JSthere


import React from 'react'
//here also you can type JS
// const MyHeading =(props) =>  {
const MyHeading =({name,text,price,pp= 22}) =>  {     //here pp=22 and in main page pp is also there then main page will be considered
    // const a = props.text
    const a = text
   //here also you can type JS         //inside this use{}for JS
  return (
    <>
    <div>{a}</div>
    <h1>{name}</h1>
    <div>MyHeading{pp}</div>         
    <MyHeading2 value={price}/>        {/*method one*/}
    </>
  )
}
const MyHeading2 =({value})=> <h2>{value}</h2>
const MyHeading3 =()=> <h3>Heading3</h3>
const MyHeading4 =()=> <h4>Heading4</h4>  //can default export only one func so we can do 2 things
export const MyHeading5 =()=> <h5>Heading3</h5> //direct export
export default MyHeading         // only default can be imported as anything
export{MyHeading3,MyHeading4};
// export default MyHeading