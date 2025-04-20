import { Button, ButtonText } from "@/components/ui/button"
import {useAuth0, Auth0Provider} from 'react-native-auth0';

export default function LoginButton() {
    const {authorize} = useAuth0();

    const onPress = async () => {
        try {
            await authorize();
        } catch (e) {
            console.log(e);
        }
    };

    return <Button onPress={onPress}>Login</Button>
}