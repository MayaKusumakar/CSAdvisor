import { Image, StyleSheet, Platform } from 'react-native';

import { HelloWave } from '@/components/HelloWave';
import ParallaxScrollView from '@/components/ParallaxScrollView';
import { ThemedText } from '@/components/ThemedText';
import { ThemedView } from '@/components/ThemedView';
import { Button, ButtonText } from "@/components/ui/button"
import LoginButton from "@/components/login/LoginButton";
import {useNavigation} from "@react-navigation/native";
import {useEffect, useState} from "react";
import { Text } from '@/components/ui/text';


export default function HomeScreen() {

   let [someState, setSomeState] = useState("Loading...");

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
          'http://localhost:8080/getEvents', {
            method: 'GET',
            headers: {
              "Content-Type": "application/json"
            },
            // body: JSON.stringify(postBody)
          }
      )
      if (response.status === 200) {
        // If status OK
        // the data returned from the API will be in JSON
          // You can extract the json from the response using this line of code
        const data = await response.json();
        // data is a json object with your data.

          setSomeState(JSON.stringify(JSON.parse(data.data)[1]));
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
        <Button size="md" variant="solid" action="primary" onPress={() => navigation.navigate('login')}>
            <ButtonText>Hello World!</ButtonText>
        </Button>
      <Text>{someState}</Text>
    </ParallaxScrollView>
  );
}

const styles = StyleSheet.create({
  titleContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 8,
  },
  stepContainer: {
    gap: 8,
    marginBottom: 8,
  },
  reactLogo: {
    height: 178,
    width: 290,
    bottom: 0,
    left: 0,
    position: 'absolute',
  },
});
