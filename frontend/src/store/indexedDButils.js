const DB_NAME = "authDB";
const DB_VERSION = 2;
let DB;

export default {
    async getDb() {
        return new Promise((resolve, reject) => {
            if (DB) {
                return resolve(DB);
            }
            let request = window.indexedDB.open(DB_NAME, DB_VERSION);

            request.onerror = e => {
                reject("Error");
                console.log(e)
            };

            request.onsuccess = e => {
                DB = e.target.result;
                resolve(DB);
            };

            request.onupgradeneeded = e => {
                console.log("onupgradeneeded called");
                let db = e.target.result;
                let objectStore = db.createObjectStore("userAuth", {
                    autoIncrement: true,
                    keyPath: "email"
                });
                console.log(objectStore)
            };
        });
    },
    async addUser(user, admin = false) {
        let db = await this.getDb();

        return new Promise(resolve => {
            let trans = db.transaction(["userAuth"], "readwrite");
            trans.oncomplete = () => {
                resolve();
            };

            let store = trans.objectStore("userAuth");
            store.openCursor().onsuccess = e => {
                let cursor = e.target.result;
                if (cursor && admin) {
                    cursor.continue();
                }
                if (cursor && !admin) {
                    if (cursor.value.email !== user.email) {
                        store.delete(cursor.value.email);
                    }
                    cursor.continue();
                }
            };
            store.put(user);
        });
    },
    async removeUser(user) {
        if (!user || !user.email) {
            return;
        }
        let db = await this.getDb();
        return new Promise(resolve => {
            let trans = db.transaction(["userAuth"], "readwrite");
            trans.oncomplete = () => {
                resolve();
            };

            let store = trans.objectStore("userAuth");
            store.delete(user.email);
            store.openCursor().onsuccess = e => {
                let cursor = e.target.result;
                if (cursor && cursor.value.access === "admin") {
                    store.delete(cursor.value.email);
                }
            };
        });
    },
    async getUser() {
        let db = await this.getDb();

        return new Promise(resolve => {
            let trans = db.transaction(["userAuth"], "readwrite");
            trans.oncomplete = () => {
                resolve(user);
            };
            let user = null;
            let store = trans.objectStore("userAuth");
            store.openCursor().onsuccess = e => {
                let cursor = e.target.result;
                if (!cursor) return;
                user = cursor.value;
                if (cursor.value.access !== "admin") {
                    cursor.continue();
                }
            };
        });
    }
};
