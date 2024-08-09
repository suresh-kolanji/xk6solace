import {Messagingservice,DirectPublish,DisconnectMessService,TerminateDirectPublish} from 'k6/x/ikeaxk6';
import { SharedArray } from 'k6/data';
const host = "localhost";
const vpn="default";
const user="admin";
const pas="admin";
export const options = {
    stages:[
 //       {duration:'20m',target:'5'},
        {duration:'5s',target:'3'},
//        {duration:'100s',target:'7'},
    ], 
  };
const connection = new Messagingservice({
	host: "localhost",
	vpn : "default",
	userName: "admin",
	pasword:"admin",
});

const data = new SharedArray('order', function () {
  // here you can open files, and then do additional processing or generate the array with data dynamically
  const f = open('./order.xml');
	const dataArray = [];
	dataArray[0]=f;
  return dataArray; // f must be an array[]
});

export default function () {
	const publish = new DirectPublish ({connection:connection,topic:"suresh",message:data[0]});
	TerminateDirectPublish(publish);
}

export function teardown(data) {
    if (__VU == 0) {
		DisconnectMessService(connection);
    }
}
