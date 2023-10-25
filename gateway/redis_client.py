import redis
import json

class RedisClient:
    def __init__(self, host='localhost', port=6379, db=0):
        self.redis_client = redis.StrictRedis(host=host, port=port, db=db)

    def add_pol_coords(self, city, resp):
        try:
            for pol in resp['policeInfo']:
                if 'confirmationNotification' not in pol:
                    self.redis_client.geoadd(city + "Police", (pol["polLong"], pol["polLat"], str(pol["polLong"]) + str(pol["polLat"]) + "KEK" + str(pol["confirmedBy"])))
                else:
                    self.redis_client.geoadd(city + "Police", (pol["polLong"], pol["polLat"], str(pol["polLong"]) + str(pol["polLat"])+ "KEK" + str(pol['confirmationNotification'])))
        except Exception as e:
            print(f"SMTH WRONG IWTH REDIS{e}")

    def add_acc_coords(self, city, resp):
        for pol in resp['accident_info']:
            if pol['confirmation_notification'] == False:
                self.redis_client.geoadd(city + "Police", (pol["pol_long"], pol["pol_lat"], pol["confirmedBy"]))
                self.expire(city+"Police", 300)


    def get_pol_values(self,city, long,lat):
        res = list()
        results = self.redis_client.georadius(city+"Police", long, lat, 1000000, unit='km', withdist=True, withcoord=True) 
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
        results = self.redis_client.georadius(city+"Accident", long, lat, 1000, unit='km', withdist=True, withcoord=True) 
        for location in results:
            res.append(json.loads({
                "confirmedBy" : location[0],
                "coordinates"  : location[2]
            }))
        return res


    def delete_pol_city_info(self, city):
        return self.redis_client.delete(city+"Police") 
    
    def delete_acc_city_info(self, city):
        return self.redis_client.delete(city+"Accident")