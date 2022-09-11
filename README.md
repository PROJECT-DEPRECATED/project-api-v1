# Project API
REST API backend service

## Feature
- Minecraft Profile
- Hanriver's info
- Home LED Light Controller
- Check current time

## Docs
<hr/>

### **/v1/hangang [GET](https://api.projecttl.net/v1/hangang)**

Success 200
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>area</td>
        <td>String</td>
        <td>
            Showing river's name
        </td>
    </tr>
    <tr>
        <td>date</td>
        <td>String</td>
        <td>
            Showing measure date
        </td>
    </tr>
    <tr>
        <td>ph</td>
        <td>String</td>
        <td>
            Showing PH Level
        </td>
    </tr>
    <tr>
        <td>respond_time</td>
        <td>String</td>
        <td>
            Showing respond time
        </td>
    </tr>
    <tr> 
        <td>status</td>
        <td>Number</td>
        <td>
            HTTP Status
        </td>
    </tr>
    <tr>
        <td>temp</td>
        <td>String</td>
        <td>
            Showing river's temperature
        </td>
    </tr>
    <tr>
        <td>time</td>
        <td>String</td>
        <td>
            Showing measure time
        </td>
    </tr>
</table>

### **/v1/hangang/:area [GET](https://api.projecttl.net/v1/hangang/tancheon)**
Parameter
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>area</td>
        <td>String</td>
        <td>
            Type area name<br/>
            - tancheon<br/>
            - jungnangcheon<br/>
            - anyang<br/>
            - seonyu<br/>
            - noryangjin<br/>
        </td>
    </tr>
</table>

Success 200
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>area</td>
        <td>String</td>
        <td>
            Showing river's name
        </td>
    </tr>
    <tr>
        <td>date</td>
        <td>String</td>
        <td>
            Showing measure date
        </td>
    </tr>
    <tr>
        <td>ph</td>
        <td>String</td>
        <td>
            Showing PH Level
        </td>
    </tr>
    <tr>
        <td>respond_time</td>
        <td>String</td>
        <td>
            Showing respond time
        </td>
    </tr>
    <tr>
        <td>status</td>
        <td>Number</td>
        <td>
            HTTP Status
        </td>
    </tr>
    <tr>
        <td>temp</td>
        <td>String</td>
        <td>
            Showing river's temperature
        </td>
    </tr>
    <tr>
        <td>time</td>
        <td>String</td>
        <td>
            Showing measure time
        </td>
    </tr>
</table>

Not Found 404
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>status</td>
        <td>Number</td>
        <td>
            Http Status
        </td>
    </tr>
    <tr>
        <td>respond_time</td>
        <td>String</td>
        <td>
            Showing respond time
        </td>
    </tr>
    <tr>
        <td>error</td>
        <td>String</td>
        <td>
            Error name
        </td>
    </tr>
</table>
<hr/>

### **/v1/led [GET]**
Headers
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>Authorization</td>
        <td>String</td>
        <td>
            Write Authentication token here
        </td>
    </tr>
</table>

Success 200
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>status</td>
        <td>Number</td>
        <td>HTTP request number</td>
    </tr>
    <tr>
        <td>type</td>
        <td>String</td>
        <td>HTTP request method type</td>
    </tr>
    <tr>
        <td>red</td>
        <td>Number</td>
        <td></td>
    </tr>
    <tr>
        <td>green</td>
        <td>Number</td>
        <td></td>
    </tr>
    <tr>
        <td>blue</td>
        <td>Number</td>
        <td></td>
    </tr>
</table>
<hr/>

### **/v1/led [POST]**
Headers
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>Authorization</td>
        <td>String</td>
        <td>
            Write Authentication token here
        </td>
    </tr>
</table>

Forms
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>red</td>
        <td>Number</td>
        <td>0~255</td>
    </tr>
    <tr>
        <td>green</td>
        <td>Number</td>
        <td>0~255</td>
    </tr>
    <tr>
        <td>blue</td>
        <td>Number</td>
        <td>0~255</td>
    </tr>
</table>

Success 200
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>status</td>
        <td>Number</td>
        <td>HTTP request number</td>
    </tr>
    <tr>
        <td>type</td>
        <td>String</td>
        <td>HTTP request method type</td>
    </tr>
    <tr>
        <td>red</td>
        <td>Number</td>
        <td></td>
    </tr>
    <tr>
        <td>green</td>
        <td>Number</td>
        <td></td>
    </tr>
    <tr>
        <td>blue</td>
        <td>Number</td>
        <td></td>
    </tr>
</table>

### **/v1/mcprofile/:username [GET](https://api.projecttl.net/v1/mcprofile/Project_TL)**
Parameter
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>username</td>
        <td>String</td>
        <td>Type minecraft player's nickname</td>
    </tr>
</table>

Success 200
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>unique_id</td>
        <td>String</td>
        <td>
            Player's uuid
        </td>
    </tr>
    <tr>
        <td>username</td>
        <td>String</td>
        <td>
            Player name
        </td>
    </tr>
    <tr>
        <td>skin_url</td>
        <td>String</td>
        <td>
            Player skin url
        </td>
    </tr>
    <tr>
        <td>respond_time</td>
        <td>String</td>
        <td>
            Showing respond time
        </td>
    </tr>
    <tr>
        <td>status</td>
        <td>Number</td>
        <td>
            HTTP Status
        </td>
    </tr>
</table>

No Content 204, Bad Request 400, Not Found 404
<table border="1">
    <th>Field</th>
    <th>Type</th>
    <th>Description</th>
    <tr>
        <td>status</td>
        <td>Number</td>
        <td>
            Http Status
        </td>
    </tr>
    <tr>
        <td>respond_time</td>
        <td>String</td>
        <td>
            Showing respond time
        </td>
    </tr>
    <tr>
        <td>error</td>
        <td>String</td>
        <td>
            Error name
        </td>
    </tr>
</table>
