package main

// base64 encode the 4 files
// I didn't want to have any external files beside the client binary
var INDEX_HTML = "PCFET0NUWVBFIGh0bWwgUFVCTElDICItLy9XM0MvL0RURCBYSFRNTCAxLjAgVHJhbnNpdGlvbmFsLy9FTiIgImh0dHA6Ly93d3cudzMub3JnL1RSL3hodG1sMS9EVEQveGh0bWwxLXRyYW5zaXRpb25hbC5kdGQiPgo8aHRtbCB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94aHRtbCIgbmctYXBwPSJjaGF0QXBwIj4KPGhlYWQ-CjxtZXRhIGh0dHAtZXF1aXY9IkNvbnRlbnQtVHlwZSIgY29udGVudD0idGV4dC9odG1sO2NoYXJzZXQ9VVRGLTgiPgo8dGl0bGU-Q2hhdCBVSTwvdGl0bGU-CjxsaW5rIHJlbD0ic3R5bGVzaGVldCIgdHlwZT0idGV4dC9jc3MiIGhyZWY9ImNoYXRhcHAuY3NzIj4KPC9oZWFkPgo8Ym9keSBuZy1jb250cm9sbGVyPSJDaGF0Q3RybCIgYmdjb2xvcj0iIzFBMjIyRiI-CjxkaXYgaWQ9ImxvZ2luZm9ybSI-CiAgICA8Zm9ybSBuZy1zdWJtaXQ9InJlZ2lzdGVyKCkiPgogICAgICAgIDxpbnB1dCBuZy1tb2RlbD0idXNlcm5hbWUiIHR5cGU9InRleHQiIGlkPSJuYW1lIiBwbGFjZWhvbGRlcj0iZW50ZXIgeW91ciBuYW1lIiBuZy1kaXNhYmxlZD0icmVnaXN0ZXJlZCIvPgogICAgICAgIDxpbnB1dCB0eXBlPSJzdWJtaXQiIHZhbHVlPSJSZWdpc3RlciIgaWQ9ImJ1dHRvbiIgbmctZGlzYWJsZWQ9InJlZ2lzdGVyZWQiPgogICAgPC9mb3JtPgo8L2Rpdj4KPGRpdiBpZD0id3JhcHBlciI-CiAgICA8ZGl2IGlkPSJtZW51Ij4KICAgICAgICA8cCBjbGFzcz0id2VsY29tZSI-V2VsY29tZSwgPGI-e3t1c2VybmFtZX19PC9iPjwvcD4KICAgICAgICA8ZGl2IHN0eWxlPSJjbGVhcjpib3RoIj48L2Rpdj4KICAgIDwvZGl2PgoKICAgIDxkaXYgc2Nyb2xsLWdsdWU-PHAgbmctcmVwZWF0PSJtZXNzYWdlIGluIG1lc3NhZ2VzIj48Yj57e21lc3NhZ2UubmFtZX19OiA8L2I-e3ttZXNzYWdlLnRleHR9fTwvcD48L2Rpdj4KCiAgICA8Zm9ybSBuZy1zdWJtaXQ9InJlcXVlc3RLZXkoKSI-CiAgICAgICAgPGlucHV0IG5nLW1vZGVsPSJ0YXJnZXR1c2VyIiB0eXBlPSJ0ZXh0IiBpZD0idGFyZ2V0IiBuZy1kaXNhYmxlZD0iISByZWdpc3RlcmVkIHx8IGNvbm5lY3RlZCIgcGxhY2Vob2xkZXI9ImNoYXQgcGFydG5lciBuYW1lIiAvPgogICAgICAgIDxpbnB1dCB0eXBlPSJzdWJtaXQiIGlkPSJidXR0b24iIHZhbHVlPSJDb25uZWN0IiBuZy1kaXNhYmxlZD0iISByZWdpc3RlcmVkIHx8IGNvbm5lY3RlZCIvPgogICAgPC9mb3JtPgogICAgPGZvcm0gbmctc3VibWl0PSJzZW5kbWVzc2FnZSgpIj4KICAgICAgICA8aW5wdXQgbmctbW9kZWw9InVzZXJtc2ciIHR5cGU9InRleHQiIGlkPSJ1c2VybXNnIiBuZy1kaXNhYmxlZD0iISByZWdpc3RlcmVkIHx8ICEgY29ubmVjdGVkIiBwbGFjZWhvbGRlcj0iY2hhdCBtZXNzYWdlIiAvPgogICAgICAgIDxpbnB1dCB0eXBlPSJzdWJtaXQiIGlkPSJidXR0b24iIHZhbHVlPSJTZW5kIiBuZy1kaXNhYmxlZD0iISByZWdpc3RlcmVkIHx8ICEgY29ubmVjdGVkIi8-CiAgICA8L2Zvcm0-CjwvZGl2PgoKPHNjcmlwdCBzcmM9Imh0dHBzOi8vYWpheC5nb29nbGVhcGlzLmNvbS9hamF4L2xpYnMvYW5ndWxhcmpzLzEuMy4wLWJldGEuMTkvYW5ndWxhci5taW4uanMiPjwvc2NyaXB0Pgo8c2NyaXB0IHNyYz0ic2Nyb2xsZ2x1ZS5qcyI-PC9zY3JpcHQ-CjxzY3JpcHQgc3JjPSJjaGF0YXBwLmpzIj48L3NjcmlwdD4KCjwvc2NyaXB0Pgo8L2JvZHk-CjwvaHRtbD4="
var SCROLLGLUE_JS = "KGZ1bmN0aW9uKGFuZ3VsYXIsIHVuZGVmaW5lZCl7CiAgICAndXNlIHN0cmljdCc7CgogICAgZnVuY3Rpb24gZmFrZU5nTW9kZWwoaW5pdFZhbHVlKXsKICAgICAgICByZXR1cm4gewogICAgICAgICAgICAkc2V0Vmlld1ZhbHVlOiBmdW5jdGlvbih2YWx1ZSl7CiAgICAgICAgICAgICAgICB0aGlzLiR2aWV3VmFsdWUgPSB2YWx1ZTsKICAgICAgICAgICAgfSwKICAgICAgICAgICAgJHZpZXdWYWx1ZTogaW5pdFZhbHVlCiAgICAgICAgfTsKICAgIH0KCiAgICBhbmd1bGFyLm1vZHVsZSgnbHVlZ2cuZGlyZWN0aXZlcycsIFtdKQogICAgLmRpcmVjdGl2ZSgnc2Nyb2xsR2x1ZScsIGZ1bmN0aW9uKCl7CiAgICAgICAgcmV0dXJuIHsKICAgICAgICAgICAgcHJpb3JpdHk6IDEsCiAgICAgICAgICAgIHJlcXVpcmU6IFsnP25nTW9kZWwnXSwKICAgICAgICAgICAgcmVzdHJpY3Q6ICdBJywKICAgICAgICAgICAgbGluazogZnVuY3Rpb24oc2NvcGUsICRlbCwgYXR0cnMsIGN0cmxzKXsKICAgICAgICAgICAgICAgIHZhciBlbCA9ICRlbFswXSwKICAgICAgICAgICAgICAgICAgICBuZ01vZGVsID0gY3RybHNbMF0gfHwgZmFrZU5nTW9kZWwodHJ1ZSk7CgogICAgICAgICAgICAgICAgZnVuY3Rpb24gc2Nyb2xsVG9Cb3R0b20oKXsKICAgICAgICAgICAgICAgICAgICBlbC5zY3JvbGxUb3AgPSBlbC5zY3JvbGxIZWlnaHQ7CiAgICAgICAgICAgICAgICB9CgogICAgICAgICAgICAgICAgZnVuY3Rpb24gc2hvdWxkQWN0aXZhdGVBdXRvU2Nyb2xsKCl7CiAgICAgICAgICAgICAgICAgICAgLy8gKyAxIGNhdGNoZXMgb2ZmIGJ5IG9uZSBlcnJvcnMgaW4gY2hyb21lCiAgICAgICAgICAgICAgICAgICAgcmV0dXJuIGVsLnNjcm9sbFRvcCArIGVsLmNsaWVudEhlaWdodCArIDEgPj0gZWwuc2Nyb2xsSGVpZ2h0OwogICAgICAgICAgICAgICAgfQoKICAgICAgICAgICAgICAgIHNjb3BlLiR3YXRjaChmdW5jdGlvbigpewogICAgICAgICAgICAgICAgICAgIGlmKG5nTW9kZWwuJHZpZXdWYWx1ZSl7CiAgICAgICAgICAgICAgICAgICAgICAgIHNjcm9sbFRvQm90dG9tKCk7CiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgfSk7CgogICAgICAgICAgICAgICAgJGVsLmJpbmQoJ3Njcm9sbCcsIGZ1bmN0aW9uKCl7CiAgICAgICAgICAgICAgICAgICAgdmFyIGFjdGl2YXRlID0gc2hvdWxkQWN0aXZhdGVBdXRvU2Nyb2xsKCk7CiAgICAgICAgICAgICAgICAgICAgaWYoYWN0aXZhdGUgIT09IG5nTW9kZWwuJHZpZXdWYWx1ZSl7CiAgICAgICAgICAgICAgICAgICAgICAgIHNjb3BlLiRhcHBseShuZ01vZGVsLiRzZXRWaWV3VmFsdWUuYmluZChuZ01vZGVsLCBhY3RpdmF0ZSkpOwogICAgICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIH0pOwogICAgICAgICAgICB9CiAgICAgICAgfTsKICAgIH0pOwp9KGFuZ3VsYXIpKTsK"
var CHATAPP_JS = "dmFyIGNoYXRBcHAgPSBhbmd1bGFyLm1vZHVsZSgnY2hhdEFwcCcsIFsnbHVlZ2cuZGlyZWN0aXZlcyddKTsKCmNoYXRBcHAuZmFjdG9yeSgnTXlTZXJ2aWNlJywgWyckcScsICckcm9vdFNjb3BlJywgZnVuY3Rpb24oJHEsICRyb290U2NvcGUpIHsKICAgIHZhciBTZXJ2aWNlID0ge307CiAgICB2YXIgY2FsbGJhY2tzID0ge307CiAgICB2YXIgY3VycmVudENhbGxiYWNrSWQgPSAwOwogICAgdmFyIHdzID0gbmV3IFdlYlNvY2tldCgid3M6Ly8xMjcuMC4wLjE6MTIzNDUvd3M_ZW5jb2Rpbmc9dGV4dCIpOwoKICAgIHdzLm9ub3BlbiA9IGZ1bmN0aW9uKCl7CiAgICAgICAgY29uc29sZS5sb2coIlNvY2tldCBoYXMgYmVlbiBvcGVuZWQhIik7CiAgICB9OwoKICAgIHdzLm9uY2xvc2UgPSBmdW5jdGlvbihldmVudCl7CiAgICAgICAgY29uc29sZS5sb2coIlNvY2tldCBoYXMgYmVlbiBjbG9zZWQhIiwgZXZlbnQpOwogICAgICAgICRyb290U2NvcGUuJGJyb2FkY2FzdCgnY2xvc2UnKQogICAgfTsKCiAgICB3cy5vbm1lc3NhZ2UgPSBmdW5jdGlvbihtZXNzYWdlKSB7CiAgICAgICAgbGlzdGVuZXIoSlNPTi5wYXJzZShtZXNzYWdlLmRhdGEpKTsKICAgIH07CgogICAgZnVuY3Rpb24gd2FpdEZvclNvY2tldENvbm5lY3Rpb24oc29ja2V0LCBjYWxsYmFjayl7CiAgICAgICAgc2V0VGltZW91dCgKICAgICAgICAgICAgZnVuY3Rpb24gKCkgewogICAgICAgICAgICAgICAgaWYgKHNvY2tldC5yZWFkeVN0YXRlID09PSAxKSB7CiAgICAgICAgICAgICAgICAgICAgY29uc29sZS5sb2coIkNvbm5lY3Rpb24gaXMgbWFkZSIpCiAgICAgICAgICAgICAgICAgICAgaWYoY2FsbGJhY2sgIT0gbnVsbCl7CiAgICAgICAgICAgICAgICAgICAgICAgIGNhbGxiYWNrKCk7CiAgICAgICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgICAgIHJldHVybjsKICAgICAgICAgICAgICAgIH0gZWxzZSB7CiAgICAgICAgICAgICAgICAgICAgY29uc29sZS5sb2coIndhaXQgZm9yIGNvbm5lY3Rpb24uLi4iKQogICAgICAgICAgICAgICAgICAgIHdhaXRGb3JTb2NrZXRDb25uZWN0aW9uKHNvY2tldCwgY2FsbGJhY2spOwogICAgICAgICAgICAgICAgfQogICAgICAgICAgICB9LAogICAgICAgIDEwKTsKICAgIH0KCiAgICBmdW5jdGlvbiBzZW5kUmVxdWVzdChyZXF1ZXN0LCB3aXRoQ2FsbGJhY2spIHsKICAgICAgICBpZiAod2l0aENhbGxiYWNrKSB7CiAgICAgICAgICAgIHZhciBkZWZlciA9ICRxLmRlZmVyKCk7CiAgICAgICAgICAgIHZhciBjYWxsYmFja0lkID0gZ2V0Q2FsbGJhY2tJZCgpOwogICAgICAgICAgICBjYWxsYmFja3NbY2FsbGJhY2tJZF0gPSB7CiAgICAgICAgICAgICAgICB0aW1lOiBuZXcgRGF0ZSgpLAogICAgICAgICAgICAgICAgY2I6ZGVmZXIKICAgICAgICAgICAgfTsKCiAgICAgICAgICAgIHJlcXVlc3QuY2FsbGJhY2tfaWQgPSBjYWxsYmFja0lkOwogICAgICAgIH0KCiAgICAgICAgY29uc29sZS5sb2coJ1NlbmRpbmcgcmVxdWVzdCcsIHJlcXVlc3QpOwoKICAgICAgICB3YWl0Rm9yU29ja2V0Q29ubmVjdGlvbih3cywgZnVuY3Rpb24oKXsKICAgICAgICAgICAgY29uc29sZS5sb2coIm1lc3NhZ2Ugc2VudCEhISIpOwogICAgICAgICAgICB3cy5zZW5kKEpTT04uc3RyaW5naWZ5KHJlcXVlc3QpKTsKICAgICAgICB9KTsKCiAgICAgICAgaWYgKHdpdGhDYWxsYmFjaykgewogICAgICAgICAgICByZXR1cm4gZGVmZXIucHJvbWlzZTsKICAgICAgICB9CiAgICB9CgogICAgZnVuY3Rpb24gbGlzdGVuZXIoZGF0YSkgewogICAgICAgIHZhciBtZXNzYWdlT2JqID0gZGF0YTsKICAgICAgICBjb25zb2xlLmxvZygiUmVjZWl2ZWQgZGF0YSBmcm9tIHdlYnNvY2tldDoiLCBtZXNzYWdlT2JqKTsKCiAgICAgICAgaWYoY2FsbGJhY2tzLmhhc093blByb3BlcnR5KG1lc3NhZ2VPYmouY2FsbGJhY2tfaWQpICYmIG1lc3NhZ2VPYmoudHlwZSAhPSAibXNnIikgewogICAgICAgICAgICBjb25zb2xlLmxvZygiUmVjZWl2ZWQgY2FsbGJhY2sgbnI6IixtZXNzYWdlT2JqLmNhbGxiYWNrX2lkKTsKICAgICAgICAgICAgY2FsbGJhY2tzW21lc3NhZ2VPYmouY2FsbGJhY2tfaWRdLmNiLnJlc29sdmUobWVzc2FnZU9iaik7CiAgICAgICAgICAgIGRlbGV0ZSBjYWxsYmFja3NbbWVzc2FnZU9iai5jYWxsYmFja19pZF07CiAgICAgICAgfQogICAgICAgIGVsc2UgewogICAgICAgICAgICAkcm9vdFNjb3BlLiRicm9hZGNhc3QoJ21zZycsIG1lc3NhZ2VPYmopCiAgICAgICAgfQogICAgfQoKICAgIGZ1bmN0aW9uIGdldENhbGxiYWNrSWQoKSB7CiAgICAgICAgY3VycmVudENhbGxiYWNrSWQgKz0gMTsKICAgICAgICBpZihjdXJyZW50Q2FsbGJhY2tJZCA-IDEwMDAwKSB7CiAgICAgICAgICAgIGN1cnJlbnRDYWxsYmFja0lkID0gMDsKICAgICAgICB9CiAgICAgICAgcmV0dXJuIGN1cnJlbnRDYWxsYmFja0lkOwogICAgfQoKICAgIFNlcnZpY2UucmVnaXN0ZXIgPSBmdW5jdGlvbih1c2VybmFtZSkgewogICAgICAgIHZhciByZXF1ZXN0ID0gewogICAgICAgICAgICBuYW1lIDogdXNlcm5hbWUsCiAgICAgICAgICAgIHR5cGUgOiAncmVnJwogICAgICAgIH0KCiAgICAgICAgcmV0dXJuIHNlbmRSZXF1ZXN0KHJlcXVlc3QsIHRydWUpOwogICAgfQoKICAgIFNlcnZpY2UuZ2V0VXNlcktleSA9IGZ1bmN0aW9uKHRhcmdldHVzZXIpIHsKICAgICAgICBpZiAodGFyZ2V0dXNlcikgewogICAgICAgICAgICB2YXIgcmVxdWVzdCA9IHsKICAgICAgICAgICAgICAgIHRhcmdldCA6IHRhcmdldHVzZXIsCiAgICAgICAgICAgICAgICB0eXBlIDogJ2NvbicKICAgICAgICAgICAgfQogICAgICAgIH0KCiAgICAgICAgcmV0dXJuIHNlbmRSZXF1ZXN0KHJlcXVlc3QsIHRydWUpOwogICAgfQoKICAgIFNlcnZpY2Uuc2VuZE1lc3NhZ2UgPSBmdW5jdGlvbihtZXNzYWdlKSB7CiAgICAgICAgdmFyIHJlcXVlc3QgPSB7CiAgICAgICAgICAgIHRleHQgOiBtZXNzYWdlLAogICAgICAgICAgICB0eXBlIDogJ21zZycKICAgICAgICB9CgogICAgICAgIHNlbmRSZXF1ZXN0KHJlcXVlc3QsIGZhbHNlKTsKICAgIH0KICAgIHJldHVybiBTZXJ2aWNlOwp9XSkKCmZ1bmN0aW9uIENoYXRDdHJsKCRzY29wZSwgJE15U2VydmljZSkgewogICAgJHNjb3BlLm1lc3NhZ2VzID0gW107CiAgICAkc2NvcGUubmFtZSA9ICIiOwogICAgJHNjb3BlLnRhcmdldCA9ICIiOwogICAgJHNjb3BlLnJlZ2lzdGVyZWQgPSBmYWxzZTsKICAgICRzY29wZS5jb25uZWN0ZWQgPSBmYWxzZTsKCiAgICAkc2NvcGUuJG9uKCdtc2cnLCBmdW5jdGlvbihldmVudCwgYXJncykgewogICAgICAgIGlmIChhcmdzLnR5cGUgPT09ICJ1c2VyIGRpc2Nvbm5lY3RlZCIpIHsKICAgICAgICAgICAgY29uc29sZS5sb2coInVzZXIgaXMgb2ZmLCB0YXJnZXQgZGlzYWJsZWQiKTsKICAgICAgICAgICAgJHNjb3BlLmNvbm5lY3RlZCA9IGZhbHNlOwogICAgICAgICAgICAkc2NvcGUudGFyZ2V0dXNlciA9ICIiOwogICAgICAgIH0KICAgIAogICAgICAgIGNvbnNvbGUubG9nKCJuZXcgbWVzc2FnZSBldmVudCIsIGFyZ3MpOwogICAgICAgICRzY29wZS5tZXNzYWdlcy5wdXNoKGFyZ3MpOwogICAgICAgICRzY29wZS4kZGlnZXN0KCk7CiAgICB9KTsKICAgIAogICAgJHNjb3BlLiRvbignY2xvc2UnLCBmdW5jdGlvbigpIHsKICAgICAgICBjb25zb2xlLmxvZygid3MgY2xvc2VkIik7CiAgICAgICAgJHNjb3BlLm1lc3NhZ2VzLnB1c2goe25hbWU6ICJicm93c2VyIiwgdGV4dDogImxvc3QgY29ubmVjdGlvbiB0byBsb2NhbCBjcnlwdG9jaGF0IHByb2dyYW1tLCByZXN0YXJ0IGl0ISJ9KTsKICAgICAgICAkc2NvcGUucmVnaXN0ZXJlZCA9IGZhbHNlOwogICAgICAgICRzY29wZS5jb25uZWN0ZWQgPSBmYWxzZTsKICAgICAgICAkc2NvcGUuJGRpZ2VzdCgpOwogICAgfSk7CgogICAgJHNjb3BlLnJlZ2lzdGVyID0gZnVuY3Rpb24oKSB7CiAgICAgICAgaWYgKCRzY29wZS51c2VybmFtZSkgewogICAgICAgICAgICAkc2NvcGUucmVnaXN0ZXJlZCA9IHRydWU7CiAgICAgICAgICAgIHZhciBjID0gJE15U2VydmljZS5yZWdpc3Rlcigkc2NvcGUudXNlcm5hbWUpOwoKICAgICAgICAgICAgYy50aGVuKGZ1bmN0aW9uKGRhdGEpIHsKICAgICAgICAgICAgICAgIGNvbnNvbGUubG9nKCJyZWdpc3RlciByZXNwb25zZTogIiwgZGF0YSk7CiAgICAgICAgICAgICAgICBpZiAoZGF0YS50eXBlID09PSAic3VjY2VzcyIpIHsKICAgICAgICAgICAgICAgICAgICBjb25zb2xlLmxvZygicmVnaXN0ZXIgcmVzcG9uc2U6ICIsIGRhdGEpOwogICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgZWxzZSB7CiAgICAgICAgICAgICAgICAgICAgJHNjb3BlLnJlZ2lzdGVyZWQgPSBmYWxzZTsKICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICRzY29wZS5tZXNzYWdlcy5wdXNoKGRhdGEpOwogICAgICAgICAgICB9LCBmdW5jdGlvbihkYXRhKSB7CiAgICAgICAgICAgICAgICBjb25zb2xlLmxvZygidGhlbiBlcnJvcjogIiwgZGF0YSk7CiAgICAgICAgICAgICAgICAkc2NvcGUucmVnaXN0ZXJlZCA9IGZhbHNlOwogICAgICAgICAgICB9LCBmdW5jdGlvbihkYXRhKSB7CiAgICAgICAgICAgICAgICBjb25zb2xlLmxvZygidGhlbiBub3RlICIsIGRhdGEpOwogICAgICAgICAgICAgICAgJHNjb3BlLnJlZ2lzdGVyZWQgPSBmYWxzZTsKICAgICAgICAgICAgfSk7CiAgICAgICAgfQogICAgfQoKICAgICRzY29wZS5yZXF1ZXN0S2V5ID0gZnVuY3Rpb24oKSB7CiAgICAgICAgaWYgKCRzY29wZS50YXJnZXR1c2VyKSB7CiAgICAgICAgICAgICRzY29wZS5jb25uZWN0ZWQgPSB0cnVlOwogICAgICAgICAgICB2YXIgYyA9ICRNeVNlcnZpY2UuZ2V0VXNlcktleSgkc2NvcGUudGFyZ2V0dXNlcik7CgogICAgICAgICAgICBjLnRoZW4oZnVuY3Rpb24oZGF0YSkgewogICAgICAgICAgICAgICAgaWYgKGRhdGEudHlwZSA9PT0gInN1Y2Nlc3MiKSB7CiAgICAgICAgICAgICAgICAgICAgY29uc29sZS5sb2coInJlcXVlc3Qga2V5IHJlc3BvbnNlOiAiLCBkYXRhKTsKICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIGVsc2UgewogICAgICAgICAgICAgICAgICAgICRzY29wZS5jb25uZWN0ZWQgPSBmYWxzZTsKICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgICRzY29wZS5tZXNzYWdlcy5wdXNoKGRhdGEpOwogICAgICAgICAgICB9LCBmdW5jdGlvbihkYXRhKSB7CiAgICAgICAgICAgICAgICBjb25zb2xlLmxvZygidGhlbiBlcnJvcjogIiwgZGF0YSk7CiAgICAgICAgICAgICAgICAkc2NvcGUuY29ubmVjdGVkID0gZmFsc2U7CiAgICAgICAgICAgIH0sIGZ1bmN0aW9uKGRhdGEpIHsKICAgICAgICAgICAgICAgIGNvbnNvbGUubG9nKCJ0aGVuIG5vdGUgIiwgZGF0YSk7CiAgICAgICAgICAgICAgICAkc2NvcGUuY29ubmVjdGVkID0gZmFsc2U7CiAgICAgICAgICAgIH0pOwogICAgICAgIH0KICAgIH0KCiAgICAkc2NvcGUuc2VuZG1lc3NhZ2UgPSBmdW5jdGlvbigpIHsKICAgICAgICBpZiAoJHNjb3BlLnVzZXJtc2cpIHsKICAgICAgICAgICAgJE15U2VydmljZS5zZW5kTWVzc2FnZSgkc2NvcGUudXNlcm1zZyk7CiAgICAgICAgICAgICRzY29wZS5tZXNzYWdlcy5wdXNoKHtuYW1lOiAkc2NvcGUudXNlcm5hbWUsIHRleHQ6ICRzY29wZS51c2VybXNnfSk7CiAgICAgICAgICAgICRzY29wZS51c2VybXNnID0gIiI7CiAgICAgICAgfQogICAgfQp9CgpjaGF0QXBwLmNvbnRyb2xsZXIoJ0NoYXRDdHJsJywgWyckc2NvcGUnLCAnTXlTZXJ2aWNlJywgQ2hhdEN0cmxdKTs="
var CHATAPP_CSS = "Ym9keSB7CiAgICBmb250OjEycHggYXJpYWw7CiAgICBjb2xvcjogIzIyMjsKICAgIHRleHQtYWxpZ246Y2VudGVyOwogICAgcGFkZGluZzozNXB4OyB9Cgpmb3JtLCBwLCBzcGFuIHsKICAgIG1hcmdpbjowOwogICAgcGFkZGluZzowOyB9CgppbnB1dCB7IGZvbnQ6MTJweCBhcmlhbDsgfQoKYSB7CiAgICBjb2xvcjojRDI0MzM1OwogICAgdGV4dC1kZWNvcmF0aW9uOm5vbmU7IH0KCiAgICBhOmhvdmVyIHsgdGV4dC1kZWNvcmF0aW9uOnVuZGVybGluZTsgfQoKI3dyYXBwZXIsICNsb2dpbmZvcm0gewogICAgbWFyZ2luOjAgYXV0bzsKICAgIHBhZGRpbmctYm90dG9tOjI1cHg7CiAgICBiYWNrZ3JvdW5kOiMwMDhFNzQ7CiAgICB3aWR0aDo1MDRweDsKICAgIGJvcmRlcjoxcHggc29saWQgI0ZGRTlBRDsgfQoKI2xvZ2luZm9ybSB7CiAgICBwYWRkaW5nLXRvcDoxOHB4OwogICAgbWFyZ2luLWJvdHRvbToxMHB4IH0KCltzY3JvbGwtZ2x1ZV0gewogICAgdGV4dC1hbGlnbjpsZWZ0OwogICAgbWFyZ2luOjAgYXV0bzsKICAgIG1hcmdpbi1ib3R0b206MjVweDsKICAgIHBhZGRpbmc6MTBweDsKICAgIGJhY2tncm91bmQ6I0IyRTA5NzsKICAgIGhlaWdodDoyNzBweDsKICAgIHdpZHRoOjQzMHB4OwogICAgYm9yZGVyOjFweCBzb2xpZCAjRkZFOUFEOwogICAgb3ZlcmZsb3c6YXV0bzsgfQoKI3VzZXJtc2csICNuYW1lLCAjdGFyZ2V0IHsKICAgIHdpZHRoOjM3MHB4OwogICAgYmFja2dyb3VuZDojQjJFMDk3OwogICAgYm9yZGVyOjFweCBzb2xpZCAjRkZFOUFEO30KCiNidXR0b24ge3dpZHRoOjgwcHg7aGVpZ2h0OjE5cHg7bWFyZ2luLWJvdHRvbToycHg7bWFyZ2luLXRvcDoycHg7CmJvcmRlcjogMXB4IHNvbGlkICNGRkU5QUQ7CmJhY2tncm91bmQtY29sb3I6ICMwMDhFNzQ7fQoKI21lbnUgeyBwYWRkaW5nOjEyLjVweCAyNXB4IDEyLjVweCAyNXB4OyB9Cgoud2VsY29tZSB7IGZsb2F0OmxlZnQ7IH0K"