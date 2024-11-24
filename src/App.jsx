import Anythin from "./components/MyHeading";
import {MyHeading3 as Renamed , MyHeading4,MyHeading5} from"./components/MyHeading";
// puprpose of arrow functio

// function MyHeading(){
//     return <div>heading one</div>
// }
// const MyHeading =()=> <h1>Heading</h1>    //it became one line code and no need to use return
function App(){
    return (                         //have to always return something
        <div>
            <Anythin name='ak' text='good going' price={2323}/>           {/*this is used as component to use other functions*/}
            <Renamed />
            <MyHeading4 />
            <MyHeading5 />
            <p>djdjdj</p>
            
        </div>
        // <div></div>    cant have two components returned 
        //to have more than one component returned you can us <></> tag or <ragment>
    )
}
export default App;   //to send whatever is coded in main page
// rcpe short cut bhai 