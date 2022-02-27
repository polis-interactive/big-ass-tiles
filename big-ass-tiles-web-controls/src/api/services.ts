import {ControlClient} from "./ControlServiceClientPb";
import {ControlRequest} from "./control_pb";

const Host = `http://${window.location.hostname}:6969`;

const Client = new ControlClient(Host, null, null);

export const TrySendControlValue = async (inputId: number, inputValue: number) => {
    try {
        let controlRequest = new ControlRequest()
        controlRequest.setInput(inputId)
        controlRequest.setValue(inputValue)
        await Client.requestControl(controlRequest, {})
    } catch (e) {
        console.error(`failed to send control request for (${inputId} ${inputValue}) with error:`, e)
    }
}
