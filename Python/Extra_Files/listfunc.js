markList=[78,17,90,90,55,95,83,95]
console.log("The marks are:"+markList+"")

markPos=markList.indexOf(9)
console.log("Index position of mark 90:"+ markPos)


markList.push(54)
console.log("After adding new marks:"+markList+"")

markList.splice(2, 0, 98)
console.log("After inserting 98 at the 2nd position:"+markList+"")

markIndex=markList.indexOf(90)
if(markIndex != -1) {
    markList.splice(markIndex, 1);
}
console.log("After removing first 90 from the list:"+markList+"")

markList.reverse()
console.log("After reversing the marks in the given list:"+markList+"")

nameList=["Marcus","James","Andrew","Bond"]
console.log("Before Sorting:"+nameList)
nameList.sort()
console.log("After Sorting:"+nameList)