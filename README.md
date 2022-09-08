# Project API
REST API backend service

## Docs
<hr/>

### **/v1/hangang [GET]**

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
        <tr>
        <td>ph</td>
        <td>String</td>
        <td>
            Showing PH Level
        </td>
    </tr>
    <tr>
        <tr>
        <td>respond_time</td>
        <td>String</td>
        <td>
            Showing get time
        </td>
    </tr>
    <tr>
        <tr>
        <td>status</td>
        <td>String</td>
        <td>
            HTTP Status
        </td>
    </tr>
    <tr>
        <tr>
        <td>temp</td>
        <td>String</td>
        <td>
            Showing river's temperature
        </td>
    </tr>
    <tr>
        <tr>
        <td>time</td>
        <td>String</td>
        <td>
            Showing measure time
        </td>
    </tr>
</table>
<hr/>

### **/v1/hangang/:area [GET]**
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
        <tr>
        <td>ph</td>
        <td>String</td>
        <td>
            Showing PH Level
        </td>
    </tr>
    <tr>
        <tr>
        <td>respond_time</td>
        <td>String</td>
        <td>
            Showing get time
        </td>
    </tr>
    <tr>
        <tr>
        <td>status</td>
        <td>String</td>
        <td>
            HTTP Status
        </td>
    </tr>
    <tr>
        <tr>
        <td>temp</td>
        <td>String</td>
        <td>
            Showing river's temperature
        </td>
    </tr>
    <tr>
        <tr>
        <td>time</td>
        <td>String</td>
        <td>
            Showing measure time
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
        <tr>
        <td>blue</td>
        <td>Number</td>
        <td></td>
    </tr>
</table>