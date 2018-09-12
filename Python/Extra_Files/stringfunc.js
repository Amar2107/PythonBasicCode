boardingCall="Good Evening, this is the final call to AI passengers for the flight AI 466 which is planned to take off at 8.40A.M."
if(boardingCall.startsWith("Good Evening"))
    console.log(boardingCall.replace("Good Evening","Good Morning"));
if(boardingCall.search("AI")>=0)
    console.log("Welcome to Air India.")
if(boardingCall.endsWith("A.M."))
    console.log("Passengers are requested to have their breakfast.");
a=boardingCall.split(" ")
for (i=0;i< a.length;i++ ){
    if(!isNaN(a[i])){
        console.log("Flight Number is specified to passengers.")
    }
}
message="Thank you all..Have a nice journey!"
console.log(message.toUpperCase())
console.log(message.toLowerCase())

