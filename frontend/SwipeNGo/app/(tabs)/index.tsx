import { Image, StyleSheet, Platform } from 'react-native';

import { HelloWave } from '@/components/HelloWave';
import ParallaxScrollView from '@/components/ParallaxScrollView';
import { ThemedText } from '@/components/ThemedText';
import { ThemedView } from '@/components/ThemedView';
import { Button, ButtonText } from "@/components/ui/button"
import LoginButton from "@/components/login/LoginButton";
import {useNavigation} from "@react-navigation/native";
import {useEffect, useState} from "react";
import { Box } from "@/components/ui/box";
import { Text } from '@/components/ui/text';


export default function HomeScreen() {

  type Event = {
    event_id: string;
    owner_id: string;
    owner_name: string;
    startTime: string;
    endTime: string;
    location: {
      latitude: number;
      longitude: number;
      address: string;
    };
    title: string;
    description: string;
    num_attending: number;
    banner_url: string;
  };
  const dummyData: Event[] = 
  [
    {
      "event_id": "abc123",
      "owner_id": "u1",
      "owner_name": "Maya",
      "startTime": "2025-04-20T10:00:00Z",
      "endTime": "2025-04-20T12:00:00Z",
      "location": {
        "latitude": 38.5449,
        "longitude": -121.7405,
        "address": "UC Davis"
      },
      "title": "Morning Coffee & Chat",
      "description": "Grab a cup and meet new people.",
      "num_attending": 12,
      "banner_url": "https://unsplash.it/400/200"
    }
  ]

   let [someState, setSomeState] = useState("Loading...");
   const [event, setEvents] = useState<Event[]>([]);

   // To make an API call, you need to have it in a useEffect react hook.
    // You can either trigger the api call on page load with no dependencies,
    // or whenever a state changes by adding the state to the list of dependencies
  useEffect(() => {
      // In the useEffect, you must declare an async function as follows to make the api call
      // Call fetch in your async function, and after your function declareation, call the function
      // Most of your logic should be in the async function
      // Consider having a state variable that tracks the function progress,
      // ie; a loading state variable that is true while the async function is working, and false when the function finishes
    const apiCall = async () => {
      const postBody = {
        description: "You want to send data to the backend with JSON",
        more_description: "You can convert javascript objects to json later with JSON.stringify(object)",
        even_more_description: "Most fetches will be POST requests also because we want to send JSON data to the backend",
        even_more_more_description: "Standard convention for this is POST requests for REST API",
        even_more_more_more_description: "Use the below snippet to make requests to the backend"
      }
      const response: Response = await fetch(
          'http://localhost:8080/ping', {
            method: 'POST',
            headers: {
              "Content-Type": "application/json"
            },
            body: JSON.stringify(postBody)
          }
      )
      if (response.status === 200) {
        // If status OK
        // the data returned from the API will be in JSON
          // You can extract the json from the response using this line of code
        const data = await response.json();
        // data is a json object with your data.

          setSomeState(JSON.stringify(data))
      } else {
        setSomeState("Error: Did not receive response status 200: " + response.status);
      }
    }
    apiCall();
  }, [])

  const navigation = useNavigation();
  return (
    
    <ParallaxScrollView
    headerBackgroundColor={{ light: '#A1CEDC', dark: '#1D3D47' }}
    headerImage={
      <Image
        source={require('@/assets/images/partial-react-logo.png')}
        style={styles.reactLogo}
      />
    }>
      
      {dummyData.map((event) => 
        <Box key={event.event_id} style = {styles.activity}>
          <Text key={event.event_id} style = {styles.title} className="text-typography-0">{event.title}</Text>
          <Image
            source={{uri: event.banner_url}}
            alt = "failed bbb"
            style = {styles.image}
            />
          
        
      </Box>
      )}
      
    </ParallaxScrollView>
  );
}

const styles = StyleSheet.create({
  activity: { top: 20, alignContent: 'center', alignItems: 'center', flex: 1, marginTop: 10, height: 640, width: 333, backgroundColor: 'pink'},
  topHalf: { flex: 1, backgroundColor: 'red', justifyContent: 'center', alignItems: 'center' },
  bottomHalf: { flex: 1, backgroundColor: 'black' },
  headerText: { textAlign: 'center', fontSize: 24, color: 'white', fontWeight: 'bold' },
  card: { margin: 10, backgroundColor: '#222', padding: 10, borderRadius: 8 },
  image: { margin: 20, width: '90%', height: 300, borderRadius: 26 },
  title: { color: 'black', fontSize: 18, fontWeight: '600', marginTop: 8 },
  desc: { color: 'gray', fontSize: 14, marginTop: 4 },
  reactLogo: {
    height: 178,
    width: 290,
    bottom: 0,
    left: 0,
    position: 'absolute',
  },
  titleContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 8,
    padding: 25,
    fontWeight: 'bold',
    fontSize: 24,
  },
  stepContainer: {
    gap: 8,
    marginBottom: 8,
  },
});
