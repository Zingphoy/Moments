async function getJsonData() {
    const url = "http://localhost:9999/test/database";
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    };
    const params = {
        method: "GET",
        cache: 'no-cache',
        headers: headers,
        mode: "cors"
    };
    try {
        let response = await fetch(url, params);
        return response.json();
    } catch (error) {
        console.log('Request Failed', error);
    }
}

function generateTableRoot(root, tableName) {
    // 根节点下插入table标签
    let tableRoot = document.createElement("table");
    tableRoot.border = "1";
    tableRoot.cellSpacing = "0";
    root.appendChild(tableRoot);

    // 写上collection名字
    let tableHead = tableRoot.createTHead();
    let row = tableHead.insertRow();
    row.innerText = tableName;
    return tableRoot;
}


function generateTableHead(tableRoot, data) {
    // 生成表头（第一行名字）
    if (data && typeof (data) !== "undefined") {
        let head = data[0];
        if (head) {
            console.log(head);
            let row = tableRoot.insertRow(-1);
            for (let attr of Object.keys(head)) {
                let cell = row.insertCell();
                cell.innerText = attr;
            }
        }
    }
}

function generateTableBody(tableRoot, row) {
    let r = tableRoot.insertRow(-1);
    for (let attr of Object.keys(row)) {
        let cell = r.insertCell();
        cell.innerText = row[attr];
    }

}

/*
  collections -> rows -> attrs
 */
function generateTables(root, jsonData) {
    for (let collection of jsonData) {
        // collection as a table
        let tableRoot = generateTableRoot(root, collection["collection"]);

        // collection name as table head
        generateTableHead(tableRoot, collection["data"]);

        // concrete collection data
        for (let row of collection["data"]) {
            // attributes as table data
            generateTableBody(tableRoot, row);
        }

        // prettify the tables
        let p1 = document.createElement("br");
        root.appendChild(p1);
        let p2 = document.createElement("br");
        root.appendChild(p2);
    }
}


function main() {
    let tableRoot = document.getElementById("mongodb");
    let jsonData = getJsonData();
    jsonData.then(data => generateTables(tableRoot, data));
}


main();
let btn = document.getElementById("refresh");
btn.onclick = function () {
    // 清楚原表格
    let tableRoot = document.getElementById("mongodb");
    tableRoot.innerHTML = "";
    main();
};







