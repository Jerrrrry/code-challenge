import http from "../http-common";
class DigService {
  
  digRecord(data) {
    return http.post("/query", data);
  }
  
  
}
export default new DigService();