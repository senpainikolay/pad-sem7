from redis.cluster import RedisCluster


class RedisClient:
    def __init__(self, host, port):
        self.port = port
        self.host = host
        self.redis_client = RedisCluster(host=self.host, port=self.port)
        print(self.redis_client.get_nodes())


    def add_pol_coords(self, city, resp):
        try:
            for pol in resp['policeInfo']:
                if 'confirmationNotification' not in pol:
                    self.redis_client.geoadd(city + "Police", (pol["polLong"], pol["polLat"], str(pol["polLong"]) + str(pol["polLat"]) + "KEK" + str(pol["confirmedBy"])))
                else:
                    self.redis_client.geoadd(city + "Police", (pol["polLong"], pol["polLat"], str(pol["polLong"]) + str(pol["polLat"])+ "KEK" + str(pol['confirmationNotification']))) 
                self.redis_client.expire(city+"Police", 30)
        except Exception as e:
            print(f"SMTH WRONG IWTH REDIS{e}")

    def add_acc_coords(self, city, resp):
        try:
            for acc in resp['accidentInfo']:
                overallStr = ""
                if 'confirmationAccidentNotification'  in acc:
                    overallStr +=  "KEK" +  "confirmationAccidentNotification"  +"KEK" + str(acc['confirmationAccidentNotification'])
                if 'confirmationPoliceNotification' in acc:
                    overallStr += "KEK" +  "confirmationPoliceNotification" + "KEK" + str(acc['confirmationPoliceNotification']) 
                if 'confirmedBy' in acc:
                    overallStr += "KEK" + "confirmedBy" + "KEK" +  str(acc['confirmedBy']) 

                self.redis_client.geoadd(city + "Accident", (acc["accidentLong"], acc["accidentLat"],
                    str(acc["accidentLong"]) + str(acc["accidentLat"])  + overallStr )) 
                self.redis_client.expire(city+"Police", 300)
        except Exception as e:
            print(f"SMTH WRONG IWTH REDIS{e}")


    def get_pol_values(self,city, long,lat):
        res = list()
        results = list()
        try:
            results = self.redis_client.georadius(city+"Police", long, lat, 100000, unit='km', withdist=True, withcoord=True) 
        except Exception as e:
            print(f"something wring with redis{e}")

        for location in results:
            kek = location[0].decode('utf-8').split("KEK")[1]
            k = "confimationNotification" 
            if kek != 'True':
                k = "confirmedBy"
            res.append({
                k : kek,
                "long"  : location[2][0],
                "lat" : location[2][1]
            })
        return  res
    
    def get_acc_values(self,city, long,lat):
        res = list()
        results = list()
        try:
            results = self.redis_client.georadius(city+"Accident", long, lat, 100000, unit='km', withdist=True, withcoord=True) 
        except Exception as e:
          
            print(f"something wring with redis{e}")

        for location in results:
            splitted = location[0].decode('utf-8').split("KEK") 
            initMap = {}
            initMap["long"] =  location[2][0]  
            initMap["lat"] = location[2][1]
            for i in range(2,len(splitted),2):
                initMap[splitted[i-1]]  =   splitted[i]
            res.append(initMap)
        return  res


    def delete_pol_city_info(self, city):
        try:
            return self.redis_client.delete(city+"Police") 
        except Exception as e:
            print(f"something wring with redis{e}")

    
    def delete_acc_city_info(self, city):
        try:
            return self.redis_client.delete(city+"Accident")
        except Exception as e:
            print(f"something wring with redis{e}")